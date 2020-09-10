# rand/guid

The benchmark compares methods of generating a globally unique string identifier (GUID).

## Random IDs

Randomly generated unique IDs are uniform and do not share attributes - e.g. [UUIDv4](https://en.wikipedia.org/wiki/Universally_unique_identifier#Version_4_(random). Each ID is entirely unpredictable, unique and non-incremental.

* [bsm/nanoid](https://github.com/bsm/nanoid)
* [google/uuid](https://github.com/google/uuid)
* [lithammer/shortuuid](https://github.com/lithammer/shortuuid)
* [matoous/go-nanoid](https://github.com/matoous/go-nanoid)

## Sortable IDs

Sortable IDs are incremental and lexicographically sortable. They usually share a common part that is host/pid specific + a current timestamp element.

* [chilts/sid](https://github.com/chilts/sid)
* [fogfish/guid](https://github.com/fogfish/guid)
* [itrabbit/rid](https://github.com/itrabbit/rid)
* [kjk/betterguid](https://github.com/kjk/betterguid)
* [rs/xid](https://github.com/rs/xid)
* [segmentio/ksuid](https://github.com/segmentio/ksuid)
* [wmentor/uniq](https://github.com/wmentor/uniq)

## Results

```
Benchmark_random/bsm/nanoid-4                      	 1604266	       743 ns/op
Benchmark_random/google/uuid-4                     	  408574	      2628 ns/op
Benchmark_random/lithammer/shortuuid-4             	  118518	     10137 ns/op
Benchmark_random/matoous/go-nanoid-4               	  368953	      2894 ns/op

Benchmark_random_parallel/bsm/nanoid-4             	 1484814	       801 ns/op
Benchmark_random_parallel/google/uuid-4            	  466240	      2346 ns/op
Benchmark_random_parallel/lithammer/shortuuid-4    	  215143	      5109 ns/op
Benchmark_random_parallel/matoous/go-nanoid-4      	  468244	      2457 ns/op

Benchmark_sorted/chilts/sid-4                      	 4327843	       276 ns/op
Benchmark_sorted/fogfish/guid-4                    	 3354352	       355 ns/op
Benchmark_sorted/itrabbit/rid-4                    	 8689556	       138 ns/op
Benchmark_sorted/kjk/betterguid-4                  	10030460	       120 ns/op
Benchmark_sorted/rs/xid-4                          	10395300	       114 ns/op
Benchmark_sorted/segmentio/ksuid-4                 	  377660	      2900 ns/op
Benchmark_sorted/wmentor/uniq-4                    	 2155467	       577 ns/op

Benchmark_sorted_parallel/chilts/sid-4             	 3509944	       334 ns/op
Benchmark_sorted_parallel/fogfish/guid-4           	 6261978	       185 ns/op
Benchmark_sorted_parallel/itrabbit/rid-4           	16508196	        65.3 ns/op
Benchmark_sorted_parallel/kjk/betterguid-4         	 6759428	       175 ns/op
Benchmark_sorted_parallel/rs/xid-4                 	19231111	        58.2 ns/op
Benchmark_sorted_parallel/segmentio/ksuid-4        	  343438	      3248 ns/op
Benchmark_sorted_parallel/wmentor/uniq-4           	 3773589	       311 ns/op
```
