# store/kv

The benchmark compares disk based, embeddable key-value stores.

- [akrylysov/pogreb](https://github.com/akrylysov/pogreb)
- [cockroachdb/pebble](https://github.com/cockroachdb/pebble)
- [dgraph-io/badger/v3](https://github.com/dgraph-io/badger/v3)
- [recoilme/pudge](https://github.com/recoilme/pudge)
- [recoilme/sniper](https://github.com/recoilme/sniper)

## Results

### no-fsync

```
Benchmark_nofsync/write/pogreb-4  	  301605	      3833 ns/op	    1272 B/op	       3 allocs/op
Benchmark_nofsync/write/badger-4  	  112146	     11213 ns/op	    1361 B/op	      36 allocs/op
Benchmark_nofsync/write/pudge-4   	  380623	      3158 ns/op	     248 B/op	      14 allocs/op
Benchmark_nofsync/write/sniper-4  	  425017	      2779 ns/op	     544 B/op	       4 allocs/op
Benchmark_nofsync/write/pebble-4  	  441164	      2560 ns/op	      11 B/op	       0 allocs/op

Benchmark_nofsync/write-batch/pogreb-4         	      22	  49249824 ns/op	     20000 keys/batch	12732430 B/op	   30453 allocs/op
Benchmark_nofsync/write-batch/badger-4         	      39	  36444416 ns/op	     20000 keys/batch	 7138155 B/op	  134653 allocs/op
Benchmark_nofsync/write-batch/pudge-4          	      34	  32113090 ns/op	     20000 keys/batch	 2484610 B/op	  140016 allocs/op
Benchmark_nofsync/write-batch/sniper-4         	      36	  33932038 ns/op	     20000 keys/batch	 5444806 B/op	   40018 allocs/op
Benchmark_nofsync/write-batch/pebble-4         	      49	  45017082 ns/op	     20000 keys/batch	 5175600 B/op	    1140 allocs/op

Benchmark_nofsync/read/pogreb-4                	 3134685	       357.0 ns/op	      64 B/op	       0 allocs/op
Benchmark_nofsync/read/badger-4                	  389431	      3163 ns/op	     580 B/op	      10 allocs/op
Benchmark_nofsync/read/pudge-4                 	 1536200	       814.5 ns/op	     112 B/op	       2 allocs/op
Benchmark_nofsync/read/sniper-4                	 1338754	       853.7 ns/op	     136 B/op	       1 allocs/op
Benchmark_nofsync/read/pebble-4                	  484526	      2196 ns/op	      64 B/op	       0 allocs/op
```
