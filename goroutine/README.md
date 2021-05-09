# Goroutine performance test

## Cases

1. Case of only append is performed.
2. Case of using goroutine to do a concurrent append.

## Result

|  Case  |  Count  |  Time  |  Memory  |  Allocate  |
| ---- | ---- | ---- | ---- |  ----  |
|  BenchmarkBasicAppend  |  7050  |  162663 ns/op  |  221184 B/op  |  1 allocs/op  |
|  BenchmarkUseGoroutineAppend  |  18  |  62157800 ns/op  |  240328 B/op  |  201 allocs/op  |

## Looking back
***We'll need to examine it in more detail.***

For simple processes like appending data to slice,  
the cost of goroutine startup and mutex.Lock becomes a bottleneck.  
Use goroutine when http requests, DB operations, or other data scanning occurs.