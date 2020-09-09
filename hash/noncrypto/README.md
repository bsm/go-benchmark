# hash/noncrypto

The benchmark compares various non-cryptographic hash methods.

* [crc64](https://golang.org/pkg/hash/crc64)
* [fnv](https://golang.org/pkg/hash/fnv)
* [cespare/xxhash](https://github.com/cespare/xxhash)
* [creachadair/cityhash](https://github.com/creachadair/cityhash)
* [dgryski/go-farm](https://github.com/dgryski/go-farm)
* [dgryski/go-metro](https://github.com/dgryski/go-metro)
* [dgryski/go-spooky](https://github.com/dgryski/go-spooky)
* [dgryski/go-stadtx](https://github.com/dgryski/go-stadtx)
* [dgryski/go-wyhash](https://github.com/dgryski/go-wyhash)
* [mmcloughlin/meow](https://github.com/mmcloughlin/meow)
* [shivakar/metrohash](https://github.com/shivakar/metrohash)
* [spaolacci/murmur3](https://github.com/spaolacci/murmur3)

## Results

```
Benchmark_64byte/crc64-4         	18614912	        59.7 ns/op	       0 B/op	       0 allocs/op
Benchmark_64byte/fnv1-4          	17491640	        66.2 ns/op	       0 B/op	       0 allocs/op
Benchmark_64byte/cespare/xxhash-4         	44026682	        25.6 ns/op	       0 B/op	       0 allocs/op
Benchmark_64byte/creachadair/cityhash-4   	24943398	        45.5 ns/op	       0 B/op	       0 allocs/op
Benchmark_64byte/dgryski/go-farm-4        	82406674	        14.0 ns/op	       0 B/op	       0 allocs/op
Benchmark_64byte/dgryski/go-metro-4       	78999768	        13.7 ns/op	       0 B/op	       0 allocs/op
Benchmark_64byte/dgryski/go-spooky-4      	21959304	        52.9 ns/op	       0 B/op	       0 allocs/op
Benchmark_64byte/dgryski/go-stadtx-4      	76584982	        15.0 ns/op	       0 B/op	       0 allocs/op
Benchmark_64byte/dgryski/go-wyhash-4      	51034880	        21.8 ns/op	       0 B/op	       0 allocs/op
Benchmark_64byte/mmcloughlin/meow-4       	 3405470	       350 ns/op	     600 B/op	       3 allocs/op
Benchmark_64byte/shivakar/metrohash-4     	49362100	        21.9 ns/op	       0 B/op	       0 allocs/op
Benchmark_64byte/spaolacci/murmur3-4      	32521585	        35.4 ns/op	       0 B/op	       0 allocs/op

Benchmark_256byte/crc64-4                 	 6708360	       184 ns/op	       0 B/op	       0 allocs/op
Benchmark_256byte/fnv1-4                  	 3720697	       327 ns/op	       0 B/op	       0 allocs/op
Benchmark_256byte/cespare/xxhash-4        	27032524	        47.8 ns/op	       0 B/op	       0 allocs/op
Benchmark_256byte/creachadair/cityhash-4  	 5348512	       206 ns/op	       0 B/op	       0 allocs/op
Benchmark_256byte/dgryski/go-farm-4       	17112612	        71.3 ns/op	       0 B/op	       0 allocs/op
Benchmark_256byte/dgryski/go-metro-4      	37797090	        35.0 ns/op	       0 B/op	       0 allocs/op
Benchmark_256byte/dgryski/go-spooky-4     	 7534816	       157 ns/op	       0 B/op	       0 allocs/op
Benchmark_256byte/dgryski/go-stadtx-4     	36302551	        33.3 ns/op	       0 B/op	       0 allocs/op
Benchmark_256byte/dgryski/go-wyhash-4     	24865210	        48.5 ns/op	       0 B/op	       0 allocs/op
Benchmark_256byte/mmcloughlin/meow-4      	 3088322	       392 ns/op	     600 B/op	       3 allocs/op
Benchmark_256byte/shivakar/metrohash-4    	25815829	        43.3 ns/op	       0 B/op	       0 allocs/op
Benchmark_256byte/spaolacci/murmur3-4     	14044209	        80.5 ns/op	       0 B/op	       0 allocs/op

Benchmark_4096byte/crc64-4                	  442494	      2537 ns/op	       0 B/op	       0 allocs/op
Benchmark_4096byte/fnv1-4                 	  222055	      5329 ns/op	       0 B/op	       0 allocs/op
Benchmark_4096byte/cespare/xxhash-4       	 3282175	       380 ns/op	       0 B/op	       0 allocs/op
Benchmark_4096byte/creachadair/cityhash-4 	  385884	      3308 ns/op	       0 B/op	       0 allocs/op
Benchmark_4096byte/dgryski/go-farm-4      	 2490466	       491 ns/op	       0 B/op	       0 allocs/op
Benchmark_4096byte/dgryski/go-metro-4     	 3716091	       334 ns/op	       0 B/op	       0 allocs/op
Benchmark_4096byte/dgryski/go-spooky-4    	 1000000	      1150 ns/op	       0 B/op	       0 allocs/op
Benchmark_4096byte/dgryski/go-stadtx-4    	 3300652	       365 ns/op	       0 B/op	       0 allocs/op
Benchmark_4096byte/dgryski/go-wyhash-4    	 3632041	       338 ns/op	       0 B/op	       0 allocs/op
Benchmark_4096byte/mmcloughlin/meow-4     	 2539258	       506 ns/op	     600 B/op	       3 allocs/op
Benchmark_4096byte/shivakar/metrohash-4   	 2337550	       479 ns/op	       0 B/op	       0 allocs/op
Benchmark_4096byte/spaolacci/murmur3-4    	 1550806	       745 ns/op	       0 B/op	       0 allocs/op
```
