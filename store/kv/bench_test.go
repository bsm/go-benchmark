package bench_test

import (
	"bytes"
	"io/ioutil"
	"math/rand"
	"os"
	"path/filepath"
	"testing"

	"github.com/akrylysov/pogreb"
	"github.com/cockroachdb/pebble"
	"github.com/dgraph-io/badger/v3"
	"github.com/recoilme/pudge"
	"github.com/recoilme/sniper"
)

const (
	keyLen  = 16
	valLen  = 128
	numKeys = 20_000
)

var value = bytes.Repeat([]byte{'x'}, valLen)

func Benchmark_nofsync(b *testing.B) {
	dir, err := ioutil.TempDir("", "go-benchmark")
	if err != nil {
		b.Fatal(err)
	}
	defer os.RemoveAll(dir)

	rnd := rand.New(rand.NewSource(33))
	keys := make([][]byte, numKeys)
	for i := 0; i < len(keys); i++ {
		key := make([]byte, keyLen)
		if _, err := rnd.Read(key); err != nil {
			b.Fatal(err)
		}
		keys[i] = key
	}

	// open pogreb
	pogrebDB, err := pogreb.Open(filepath.Join(dir, "pogreb"), nil)
	if err != nil {
		b.Fatal(err)
	}
	defer pogrebDB.Close()

	// open badger
	badgerOpt := badger.DefaultOptions(filepath.Join(dir, "badger")).
		WithLoggingLevel(badger.ERROR).
		WithSyncWrites(false)
	badgerDB, err := badger.Open(badgerOpt)
	if err != nil {
		b.Fatal(err)
	}
	defer badgerDB.Close()

	// open pudge
	pudgeDB, err := pudge.Open(filepath.Join(dir, "pudge"), &pudge.Config{SyncInterval: 0})
	if err != nil {
		b.Fatal(err)
	}
	defer pudgeDB.Close()

	// open sniper
	sniperDB, err := sniper.Open(sniper.Dir(filepath.Join(dir, "sniper")))
	if err != nil {
		b.Fatal(err)
	}
	defer sniperDB.Close()

	// open pebble
	pebbleDB, err := pebble.Open(filepath.Join(dir, "pebble"), nil)
	if err != nil {
		b.Fatal(err)
	}
	defer pebbleDB.Close()

	b.Run("write", func(b *testing.B) {
		benchWrite(b, "pogreb", keys, func(key []byte) error { return pogrebDB.Put(key, value) })
		benchWrite(b, "badger", keys, func(key []byte) error {
			return badgerDB.Update(func(txn *badger.Txn) error {
				return txn.Set(key, value)
			})
		})
		benchWrite(b, "pudge", keys, func(key []byte) error { return pudgeDB.Set(key, value) })
		benchWrite(b, "sniper", keys, func(key []byte) error { return sniperDB.Set(key, value, 0) })
		benchWrite(b, "pebble", keys, func(key []byte) error { return pebbleDB.Set(key, value, pebble.NoSync) })
	})

	b.Run("write-batch", func(b *testing.B) {
		benchWriteBatch(b, "pogreb", keys, func() error {
			for i := 0; i < len(keys); i += 2 {
				if err := pogrebDB.Put(keys[i], value); err != nil {
					return err
				}
			}
			return pogrebDB.Sync()
		})
		benchWriteBatch(b, "badger", keys, func() error {
			b := badgerDB.NewWriteBatch()
			defer b.Cancel()

			for i := 0; i < len(keys); i += 2 {
				if err := b.Set(keys[i], value); err != nil {
					return err
				}
			}
			if err := b.Flush(); err != nil {
				return err
			}
			return badgerDB.Sync()
		})
		benchWriteBatch(b, "pudge", keys, func() error {
			for i := 0; i < len(keys); i += 2 {
				if err := pudgeDB.Set(keys[i], value); err != nil {
					return err
				}
			}
			return nil
		})
		benchWriteBatch(b, "sniper", keys, func() error {
			for i := 0; i < len(keys); i += 2 {
				if err := sniperDB.Set(keys[i], value, 0); err != nil {
					return err
				}
			}
			return nil
		})
		benchWriteBatch(b, "pebble", keys, func() error {
			b := pebbleDB.NewBatch()
			defer b.Close()

			for i := 0; i < len(keys); i += 2 {
				if err := b.Set(keys[i], value, nil); err != nil {
					return err
				}
			}
			return b.Commit(pebble.Sync)
		})
	})

	b.Run("read", func(b *testing.B) {
		benchRead(b, "pogreb", keys, func(key []byte) ([]byte, error) { return pogrebDB.Get(key) })
		benchRead(b, "badger", keys, func(key []byte) ([]byte, error) {
			var val []byte
			err := badgerDB.View(func(txn *badger.Txn) error {
				item, err := txn.Get(key)
				if err != nil {
					return err
				}
				val, err = item.ValueCopy(nil)
				return err
			})
			if err == badger.ErrKeyNotFound {
				return nil, nil
			}
			return val, err
		})
		benchRead(b, "pudge", keys, func(key []byte) ([]byte, error) {
			var val []byte
			err := pudgeDB.Get(key, &val)
			if err == pudge.ErrKeyNotFound {
				return nil, nil
			}
			return val, err
		})
		benchRead(b, "sniper", keys, func(key []byte) ([]byte, error) {
			val, err := sniperDB.Get(key)
			if err == sniper.ErrNotFound {
				return nil, nil
			}
			return val, err
		})
		benchRead(b, "pebble", keys, func(key []byte) ([]byte, error) {
			val, closer, err := pebbleDB.Get(key)
			if err == pebble.ErrNotFound {
				return nil, nil
			} else if err != nil {
				return nil, err
			}

			clone := make([]byte, len(val))
			copy(clone, val)
			if err := closer.Close(); err != nil {
				return nil, err
			}
			return clone, nil
		})
	})
}

func benchWrite(b *testing.B, name string, keys [][]byte, fn func([]byte) error) {
	b.Helper()

	b.Run(name, func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			pos := (2 * i) % len(keys)
			key := keys[pos]
			if err := fn(key); err != nil {
				b.Fatal(err)
			}
		}
	})
}

func benchWriteBatch(b *testing.B, name string, keys [][]byte, fn func() error) {
	b.Helper()

	b.Run(name, func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			if err := fn(); err != nil {
				b.Fatal(err)
			}
		}
		b.ReportMetric(float64(numKeys), "keys/batch")
	})
}

func benchRead(b *testing.B, name string, keys [][]byte, fn func([]byte) ([]byte, error)) {
	b.Helper()

	b.Run(name, func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			key := keys[i%len(keys)]
			val, err := fn(key)
			if err != nil {
				b.Fatal(err)
			}
			if i%2 == 0 {
				if !bytes.Equal(val, value) {
					b.Fatalf("expected %v, got %v", value, val)
				}
			} else if val != nil {
				b.Fatalf("expected nil, got %v", val)
			}
		}
	})
}
