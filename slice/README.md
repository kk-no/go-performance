# Slice performance test

## Cases

1. Cases where no allocation is made during initialization.
2. The case of using make to allocate and append during initialization.
3. The case of using make to allocate and Index to store during initialization.

## Result

|  Case  |  Count  |  Time  |  Memory  |  Allocate  |
| ---- | ---- | ---- | ---- |  ----  |
|  BenchmarkNoMakeAppend  |  4622  |  259222 ns/op  |  1103868 B/op  |  28 allocs/op  |
|  BenchmarkUseMakeAppend  |  7344  |  161272 ns/op  |  221184 B/op  |  1 allocs/op  |
|  BenchmarkUseMakeIndex  |  13136  |  91398 ns/op  |  221184 B/op  |  1 allocs/op  |

## Looking back

Use make as much as possible when initializing slice.  
If you can use Index, it will be faster, but be aware that the length will be fixed during initialization.