# http/json

The benchmark compares methods of generating a JSON response in Go with different payloads.

1. encode directly: `json.NewEncoder(w).Encode(payload)`
2. use a `bufio.Writer` to buffer the output (with and without pooling)
3. first encode to a `*bytes.Buffer` then use `buf.WriteTo(w)` (with and without pooling)

## Results

```
Benchmark/small_payload/direct-4  	   29007	     40616 ns/op	    1476 B/op	      14 allocs/op
Benchmark/small_payload/bufio-4   	   29451	     40354 ns/op	    5647 B/op	      16 allocs/op
Benchmark/small_payload/bufio_with_pooling-4         	   29834	     39225 ns/op	    1473 B/op	      14 allocs/op
Benchmark/small_payload/buffer-4                     	   29067	     41536 ns/op	    1583 B/op	      16 allocs/op
Benchmark/small_payload/buffer_with_pooling-4        	   28321	     47815 ns/op	    1472 B/op	      14 allocs/op

Benchmark/large_payload/direct-4                     	    6988	    166522 ns/op	   73983 B/op	      19 allocs/op
Benchmark/large_payload/bufio-4                      	    6586	    162004 ns/op	   78083 B/op	      21 allocs/op
Benchmark/large_payload/bufio_with_pooling-4         	    6872	    158594 ns/op	   73880 B/op	      19 allocs/op
Benchmark/large_payload/buffer-4                     	    5959	    181169 ns/op	  116799 B/op	      22 allocs/op
Benchmark/large_payload/buffer_with_pooling-4        	    6915	    163763 ns/op	   76207 B/op	      19 allocs/op
```
