# http/json

The benchmark compares methods of generating a JSON response in Go with different payloads.

1. encode directly: `json.NewEncoder(w).Encode(payload)`
2. use a `bufio.Writer` to buffer the output (with and without pooling)
3. first encode to a `*bytes.Buffer` then use `buf.WriteTo(w)` (with and without pooling)

## Results

```
Benchmark/small_payload/direct-4         	   28822	     40302 ns/op	    1477 B/op	      14 allocs/op
Benchmark/small_payload/bufio-4          	   28213	     41653 ns/op	    5648 B/op	      16 allocs/op
Benchmark/small_payload/bufio_pool-4     	   29620	     46394 ns/op	    1474 B/op	      14 allocs/op
Benchmark/small_payload/buffer-4         	   23004	     62572 ns/op	    1578 B/op	      15 allocs/op
Benchmark/small_payload/buffer_pool-4    	   29420	     39745 ns/op	    1473 B/op	      14 allocs/op

Benchmark/large_payload/direct-4         	    6552	    157061 ns/op	   74673 B/op	      19 allocs/op
Benchmark/large_payload/bufio-4          	    6696	    161834 ns/op	   78287 B/op	      21 allocs/op
Benchmark/large_payload/bufio_pool-4     	    7094	    156547 ns/op	   74394 B/op	      19 allocs/op
Benchmark/large_payload/buffer-4         	    7082	    168776 ns/op	  116866 B/op	      22 allocs/op
Benchmark/large_payload/buffer_pool-4    	    7168	    160605 ns/op	   76104 B/op	      19 allocs/op
```
