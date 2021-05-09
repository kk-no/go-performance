# I/O performance test

## Cases

1. Case of io.ReadAll io.Reader and writing to http.ResponseWriter.
2. Case of io.Copy io.Reader and writing to http.ResponseWriter.

## Result

|  Case  |  Count  |  Time  |  Memory  |  Allocate  |
| ---- | ---- | ---- | ---- |  ----  |
|  BenchmarkIOReadAll  |  7928  |  148983 ns/op  |  1325538 B/op  |  32 allocs/op  |
|  BenchmarkIOCopy  |  45483  |  25842 ns/op  |  222176 B/op  |  10 allocs/op  |

## Looking back

When handling Streams such as io.Reader,  
it is faster to write directly with io.Reader.