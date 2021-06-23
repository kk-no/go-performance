# Json performance test

## Cases

1. The case of reading the http response with io.ReadAll and Unmarshal.
2. The case of reading the http response with json.Decoder.

## Result

|  Case  |  Count  |  Time  |  Memory  |  Allocate  |
| ---- | ---- | ---- | ---- |  ----  |
|  BenchmarkJSONUnmarshal-8  |  48  |  24710249 ns/op  |  46016960 B/op  |  171 allocs/op  |
|  BenchmarkJSONDecode-8  |  54  |  22375486 ns/op  |  23031417 B/op  |  190 allocs/op  |

## Looking back

If you use ReadAll, it will use extra memory.  
io.Reader streams can be passed directly to the Decoder to reduce memory usage.