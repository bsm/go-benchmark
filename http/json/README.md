# http/json

The benchmark compares methods of generating a JSON response in Go with different payloads.

1. encode directly: `json.NewEncoder(w).Encode(payload)`
2. use a `bufio.Writer` to buffer the output (with and without pooling)
3. first encode to a `*bytes.Buffer` then use `buf.WriteTo(w)` (with and without pooling)

## Results

```
Benchmark/small_payload/direct-4  	    7532	    153040 ns/op	   16311 B/op	     112 allocs/op
Benchmark/small_payload/bufio-4   	   19650	     60103 ns/op	    8229 B/op	      51 allocs/op
Benchmark/small_payload/bufio_with_pooling-4         	   20138	     58468 ns/op	    4067 B/op	      49 allocs/op
Benchmark/small_payload/buffer-4                     	    7131	    155010 ns/op	   16414 B/op	     114 allocs/op
Benchmark/small_payload/buffer_with_pooling-4        	    7384	    155175 ns/op	   16307 B/op	     112 allocs/op

Benchmark/large_payload/direct-4                     	    4831	    262378 ns/op	   16658 B/op	     116 allocs/op
Benchmark/large_payload/bufio-4                      	    4467	    248376 ns/op	   21201 B/op	     118 allocs/op
Benchmark/large_payload/bufio_with_pooling-4         	    4842	    246278 ns/op	   16729 B/op	     116 allocs/op
Benchmark/large_payload/buffer-4                     	    4442	    265030 ns/op	   57771 B/op	     119 allocs/op
Benchmark/large_payload/buffer_with_pooling-4        	    4538	    256824 ns/op	   16717 B/op	     116 allocs/op
```
