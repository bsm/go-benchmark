# sketch/histogram

The benchmark compares various histogram sketches using various
data distributions and quantile ranges.

- [ajwerner/tdigest](https://github.com/ajwerner/tdigest)
- [beorn7/perks/quantile](https://github.com/beorn7/perks)
- [bsm/histogram](https://github.com/bsm/histogram)
- [bsm/tdigest](https://github.com/bsm/tdigest)
- [caio/go-tdigest](https://github.com/caio/go-tdigest)
- [influxdata/tdigest](https://github.com/influxdata/tdigest)
- [segmentio/tdigest](https://github.com/segmentio/tdigest)
- [signalfx/tdigest](https://github.com/signalfx/tdigest)
- [spenczar/tdigest](https://github.com/spenczar/tdigest)
- [willbeason/tdigest](https://github.com/willbeason/tdigest)

## Results

```
BenchmarkAdd/ajwerner-4         	10786482	       109.3 ns/op
BenchmarkAdd/beorn7-4           	 3035592	       518.0 ns/op
BenchmarkAdd/bsmhisto-4         	 5032197	       239.9 ns/op
BenchmarkAdd/bsmdigest-4        	 9504627	       124.4 ns/op
BenchmarkAdd/caio-4             	 4358706	       287.2 ns/op
BenchmarkAdd/influx-4           	 9610213	       122.2 ns/op
BenchmarkAdd/spenczar-4         	 1000000	      2007 ns/op
BenchmarkAdd/segmentio-4        	 1000000	      2032 ns/op
BenchmarkAdd/signalfx-4         	12755361	        87.88 ns/op
BenchmarkAdd/willbeason-4       	22241617	        53.95 ns/op

BenchmarkQuery/uniform/ajwerner-4         	26177332	        44.06 ns/op
BenchmarkQuery/uniform/beorn7-4           	  292666	      4014 ns/op
BenchmarkQuery/uniform/bsmhisto-4         	 8047804	       147.3 ns/op
BenchmarkQuery/uniform/bsmdigest-4        	19981726	        56.01 ns/op
BenchmarkQuery/uniform/caio-4             	 1374925	       902.1 ns/op
BenchmarkQuery/uniform/influx-4           	21801067	        53.74 ns/op
BenchmarkQuery/uniform/spenczar-4         	11014292	       104.5 ns/op
BenchmarkQuery/uniform/segmentio-4        	12035796	        98.49 ns/op
BenchmarkQuery/uniform/signalfx-4         	33722883	        35.10 ns/op
BenchmarkQuery/uniform/willbeason-4       	14579332	        81.52 ns/op

BenchmarkQuery/lognorm/ajwerner-4         	21856186	        46.53 ns/op
BenchmarkQuery/lognorm/beorn7-4           	  283106	      3999 ns/op
BenchmarkQuery/lognorm/bsmhisto-4         	28575897	        40.59 ns/op
BenchmarkQuery/lognorm/bsmdigest-4        	21211870	        55.63 ns/op
BenchmarkQuery/lognorm/caio-4             	 1401163	       859.5 ns/op
BenchmarkQuery/lognorm/influx-4           	22310814	        52.57 ns/op
BenchmarkQuery/lognorm/spenczar-4         	11756181	       100.4 ns/op
BenchmarkQuery/lognorm/segmentio-4        	12499549	        95.77 ns/op
BenchmarkQuery/lognorm/signalfx-4         	33322134	        34.78 ns/op
BenchmarkQuery/lognorm/willbeason-4       	13760164	        86.80 ns/op

BenchmarkPrecision/uniform/p10/ajwerner-4 	1000000000	         1.145 %err
BenchmarkPrecision/uniform/p10/beorn7-4   	1000000000	         0.4974 %err
BenchmarkPrecision/uniform/p10/bsmhisto-4 	1000000000	         0.05316 %err
BenchmarkPrecision/uniform/p10/bsmdigest-4         	1000000000	         0.3194 %err
BenchmarkPrecision/uniform/p10/caio-4              	1000000000	         0.03112 %err
BenchmarkPrecision/uniform/p10/influx-4            	1000000000	         0.3194 %err
BenchmarkPrecision/uniform/p10/spenczar-4          	1000000000	         0.09332 %err
BenchmarkPrecision/uniform/p10/segmentio-4         	1000000000	         0.09332 %err
BenchmarkPrecision/uniform/p10/signalfx-4          	1000000000	         1.452 %err
BenchmarkPrecision/uniform/p10/willbeason-4        	1000000000	         3.636 %err

BenchmarkPrecision/uniform/p25/ajwerner-4          	1000000000	         0.4195 %err
BenchmarkPrecision/uniform/p25/beorn7-4            	1000000000	         0.7055 %err
BenchmarkPrecision/uniform/p25/bsmhisto-4          	1000000000	         0.04836 %err
BenchmarkPrecision/uniform/p25/bsmdigest-4         	1000000000	         0.006783 %err
BenchmarkPrecision/uniform/p25/caio-4              	1000000000	         0.02134 %err
BenchmarkPrecision/uniform/p25/influx-4            	1000000000	         0.006783 %err
BenchmarkPrecision/uniform/p25/spenczar-4          	1000000000	         1.491 %err
BenchmarkPrecision/uniform/p25/segmentio-4         	1000000000	         1.491 %err
BenchmarkPrecision/uniform/p25/signalfx-4          	1000000000	         0.7537 %err
BenchmarkPrecision/uniform/p25/willbeason-4        	1000000000	         2.444 %err

BenchmarkPrecision/uniform/p50/ajwerner-4          	1000000000	         0.1634 %err
BenchmarkPrecision/uniform/p50/beorn7-4            	1000000000	         0.6053 %err
BenchmarkPrecision/uniform/p50/bsmhisto-4          	1000000000	         0.002993 %err
BenchmarkPrecision/uniform/p50/bsmdigest-4         	1000000000	         0.01779 %err
BenchmarkPrecision/uniform/p50/caio-4              	1000000000	         0.002020 %err
BenchmarkPrecision/uniform/p50/influx-4            	1000000000	         0.01779 %err
BenchmarkPrecision/uniform/p50/spenczar-4          	1000000000	         2.633 %err
BenchmarkPrecision/uniform/p50/segmentio-4         	1000000000	         2.633 %err
BenchmarkPrecision/uniform/p50/signalfx-4          	1000000000	         0.08240 %err
BenchmarkPrecision/uniform/p50/willbeason-4        	1000000000	         0.008872 %err

BenchmarkPrecision/uniform/p75/ajwerner-4          	1000000000	         0.1932 %err
BenchmarkPrecision/uniform/p75/beorn7-4            	1000000000	         0.3061 %err
BenchmarkPrecision/uniform/p75/bsmhisto-4          	1000000000	         0.0000608 %err
BenchmarkPrecision/uniform/p75/bsmdigest-4         	1000000000	         0.01543 %err
BenchmarkPrecision/uniform/p75/caio-4              	1000000000	         0.003496 %err
BenchmarkPrecision/uniform/p75/influx-4            	1000000000	         0.01543 %err
BenchmarkPrecision/uniform/p75/spenczar-4          	1000000000	         0.3139 %err
BenchmarkPrecision/uniform/p75/segmentio-4         	1000000000	         0.3139 %err
BenchmarkPrecision/uniform/p75/signalfx-4          	1000000000	         0.4135 %err
BenchmarkPrecision/uniform/p75/willbeason-4        	1000000000	         0.8020 %err

BenchmarkPrecision/uniform/p90/ajwerner-4          	1000000000	         0.1270 %err
BenchmarkPrecision/uniform/p90/beorn7-4            	1000000000	         0.05536 %err
BenchmarkPrecision/uniform/p90/bsmhisto-4          	1000000000	         0.002342 %err
BenchmarkPrecision/uniform/p90/bsmdigest-4         	1000000000	         0.01965 %err
BenchmarkPrecision/uniform/p90/caio-4              	1000000000	         0.003163 %err
BenchmarkPrecision/uniform/p90/influx-4            	1000000000	         0.01965 %err
BenchmarkPrecision/uniform/p90/spenczar-4          	1000000000	         0.02129 %err
BenchmarkPrecision/uniform/p90/segmentio-4         	1000000000	         0.02129 %err
BenchmarkPrecision/uniform/p90/signalfx-4          	1000000000	         0.2118 %err
BenchmarkPrecision/uniform/p90/willbeason-4        	1000000000	         0.5073 %err

BenchmarkPrecision/uniform/p95/ajwerner-4          	1000000000	         0.03827 %err
BenchmarkPrecision/uniform/p95/beorn7-4            	1000000000	         0.02508 %err
BenchmarkPrecision/uniform/p95/bsmhisto-4          	1000000000	         0.01078 %err
BenchmarkPrecision/uniform/p95/bsmdigest-4         	1000000000	         0.01070 %err
BenchmarkPrecision/uniform/p95/caio-4              	1000000000	         0.001481 %err
BenchmarkPrecision/uniform/p95/influx-4            	1000000000	         0.01070 %err
BenchmarkPrecision/uniform/p95/spenczar-4          	1000000000	         0.1316 %err
BenchmarkPrecision/uniform/p95/segmentio-4         	1000000000	         0.1316 %err
BenchmarkPrecision/uniform/p95/signalfx-4          	1000000000	         0.07749 %err
BenchmarkPrecision/uniform/p95/willbeason-4        	1000000000	         0.3256 %err

BenchmarkPrecision/uniform/p99/ajwerner-4          	1000000000	         0.01908 %err
BenchmarkPrecision/uniform/p99/beorn7-4            	1000000000	         0.0008529 %err
BenchmarkPrecision/uniform/p99/bsmhisto-4          	1000000000	         0.02582 %err
BenchmarkPrecision/uniform/p99/bsmdigest-4         	1000000000	         0.02018 %err
BenchmarkPrecision/uniform/p99/caio-4              	1000000000	         0.002248 %err
BenchmarkPrecision/uniform/p99/influx-4            	1000000000	         0.02018 %err
BenchmarkPrecision/uniform/p99/spenczar-4          	1000000000	         0.1424 %err
BenchmarkPrecision/uniform/p99/segmentio-4         	1000000000	         0.1424 %err
BenchmarkPrecision/uniform/p99/signalfx-4          	1000000000	         0.02659 %err
BenchmarkPrecision/uniform/p99/willbeason-4        	1000000000	         0.1210 %err

BenchmarkPrecision/lognorm/p10/ajwerner-4          	1000000000	         0.6280 %err
BenchmarkPrecision/lognorm/p10/beorn7-4            	1000000000	         0.9781 %err
BenchmarkPrecision/lognorm/p10/bsmhisto-4          	1000000000	         1.558 %err
BenchmarkPrecision/lognorm/p10/bsmdigest-4         	1000000000	         0.06188 %err
BenchmarkPrecision/lognorm/p10/caio-4              	1000000000	         0.006783 %err
BenchmarkPrecision/lognorm/p10/influx-4            	1000000000	         0.06188 %err
BenchmarkPrecision/lognorm/p10/spenczar-4          	1000000000	         1.814 %err
BenchmarkPrecision/lognorm/p10/segmentio-4         	1000000000	         1.814 %err
BenchmarkPrecision/lognorm/p10/signalfx-4          	1000000000	         8.871 %err
BenchmarkPrecision/lognorm/p10/willbeason-4        	1000000000	         2.772 %err

BenchmarkPrecision/lognorm/p25/ajwerner-4          	1000000000	         0.5219 %err
BenchmarkPrecision/lognorm/p25/beorn7-4            	1000000000	         0.4641 %err
BenchmarkPrecision/lognorm/p25/bsmhisto-4          	1000000000	         0.4477 %err
BenchmarkPrecision/lognorm/p25/bsmdigest-4         	1000000000	         0.02384 %err
BenchmarkPrecision/lognorm/p25/caio-4              	1000000000	         0.01225 %err
BenchmarkPrecision/lognorm/p25/influx-4            	1000000000	         0.02384 %err
BenchmarkPrecision/lognorm/p25/spenczar-4          	1000000000	         0.6959 %err
BenchmarkPrecision/lognorm/p25/segmentio-4         	1000000000	         0.6959 %err
BenchmarkPrecision/lognorm/p25/signalfx-4          	1000000000	         2.066 %err
BenchmarkPrecision/lognorm/p25/willbeason-4        	1000000000	         1.692 %err

BenchmarkPrecision/lognorm/p50/ajwerner-4          	1000000000	         0.2640 %err
BenchmarkPrecision/lognorm/p50/beorn7-4            	1000000000	         0.2063 %err
BenchmarkPrecision/lognorm/p50/bsmhisto-4          	1000000000	         0.01073 %err
BenchmarkPrecision/lognorm/p50/bsmdigest-4         	1000000000	         0.05651 %err
BenchmarkPrecision/lognorm/p50/caio-4              	1000000000	         0.01454 %err
BenchmarkPrecision/lognorm/p50/influx-4            	1000000000	         0.05651 %err
BenchmarkPrecision/lognorm/p50/spenczar-4          	1000000000	         5.006 %err
BenchmarkPrecision/lognorm/p50/segmentio-4         	1000000000	         5.006 %err
BenchmarkPrecision/lognorm/p50/signalfx-4          	1000000000	         3.613 %err
BenchmarkPrecision/lognorm/p50/willbeason-4        	1000000000	         1.520 %err

BenchmarkPrecision/lognorm/p75/ajwerner-4          	1000000000	         0.1894 %err
BenchmarkPrecision/lognorm/p75/beorn7-4            	1000000000	         0.6658 %err
BenchmarkPrecision/lognorm/p75/bsmhisto-4          	1000000000	         0.2697 %err
BenchmarkPrecision/lognorm/p75/bsmdigest-4         	1000000000	         0.02679 %err
BenchmarkPrecision/lognorm/p75/caio-4              	1000000000	         0.003314 %err
BenchmarkPrecision/lognorm/p75/influx-4            	1000000000	         0.02679 %err
BenchmarkPrecision/lognorm/p75/spenczar-4          	1000000000	         0.7522 %err
BenchmarkPrecision/lognorm/p75/segmentio-4         	1000000000	         0.7522 %err
BenchmarkPrecision/lognorm/p75/signalfx-4          	1000000000	         1.175 %err
BenchmarkPrecision/lognorm/p75/willbeason-4        	1000000000	         0.5168 %err

BenchmarkPrecision/lognorm/p90/ajwerner-4          	1000000000	         0.3764 %err
BenchmarkPrecision/lognorm/p90/beorn7-4            	1000000000	         0.4636 %err
BenchmarkPrecision/lognorm/p90/bsmhisto-4          	1000000000	         0.1727 %err
BenchmarkPrecision/lognorm/p90/bsmdigest-4         	1000000000	         0.1441 %err
BenchmarkPrecision/lognorm/p90/caio-4              	1000000000	         0.03852 %err
BenchmarkPrecision/lognorm/p90/influx-4            	1000000000	         0.1441 %err
BenchmarkPrecision/lognorm/p90/spenczar-4          	1000000000	         1.757 %err
BenchmarkPrecision/lognorm/p90/segmentio-4         	1000000000	         1.757 %err
BenchmarkPrecision/lognorm/p90/signalfx-4          	1000000000	         0.8455 %err
BenchmarkPrecision/lognorm/p90/willbeason-4        	1000000000	         0.2821 %err

BenchmarkPrecision/lognorm/p95/ajwerner-4          	1000000000	         0.2281 %err
BenchmarkPrecision/lognorm/p95/beorn7-4            	1000000000	         0.3204 %err
BenchmarkPrecision/lognorm/p95/bsmhisto-4          	1000000000	         0.07680 %err
BenchmarkPrecision/lognorm/p95/bsmdigest-4         	1000000000	         0.003234 %err
BenchmarkPrecision/lognorm/p95/caio-4              	1000000000	         0.01142 %err
BenchmarkPrecision/lognorm/p95/influx-4            	1000000000	         0.003234 %err
BenchmarkPrecision/lognorm/p95/spenczar-4          	1000000000	         1.335 %err
BenchmarkPrecision/lognorm/p95/segmentio-4         	1000000000	         1.335 %err
BenchmarkPrecision/lognorm/p95/signalfx-4          	1000000000	         0.3462 %err
BenchmarkPrecision/lognorm/p95/willbeason-4        	1000000000	         0.9458 %err

BenchmarkPrecision/lognorm/p99/ajwerner-4          	1000000000	         0.2574 %err
BenchmarkPrecision/lognorm/p99/beorn7-4            	1000000000	         0.1536 %err
BenchmarkPrecision/lognorm/p99/bsmhisto-4          	1000000000	         0.1433 %err
BenchmarkPrecision/lognorm/p99/bsmdigest-4         	1000000000	         0.04964 %err
BenchmarkPrecision/lognorm/p99/caio-4              	1000000000	         0.09018 %err
BenchmarkPrecision/lognorm/p99/influx-4            	1000000000	         0.04964 %err
BenchmarkPrecision/lognorm/p99/spenczar-4          	1000000000	         0.6877 %err
BenchmarkPrecision/lognorm/p99/segmentio-4         	1000000000	         0.6877 %err
BenchmarkPrecision/lognorm/p99/signalfx-4          	1000000000	         0.2522 %err
BenchmarkPrecision/lognorm/p99/willbeason-4        	1000000000	         1.028 %err
```
