# Map (hash/dict) vs Slice(array/list) Benchmark

Benchmarks are made with the algorithm for calculating the frequency of characters in the text using both data structure types.
Source text: 1 chapter of 1984 Orwell (~35KB)



## Go

```shell
$ go version
go version go1.20.2 linux/amd64

$ go test -bench=. -benchmem -count=5
goos: linux
goarch: amd64
pkg: maptests
cpu: Intel(R) Core(TM) i7-1065G7 CPU @ 1.30GHz
BenchmarkMap-4              2012            564460 ns/op            5497 B/op          4 allocs/op
BenchmarkMap-4              2130            536793 ns/op            5495 B/op          4 allocs/op
BenchmarkMap-4              2193            539073 ns/op            5493 B/op          4 allocs/op
BenchmarkMap-4              2170            543566 ns/op            5494 B/op          4 allocs/op
BenchmarkMap-4              2097            544182 ns/op            5495 B/op          4 allocs/op
BenchmarkSlice-4           34066             34256 ns/op            1050 B/op          2 allocs/op
BenchmarkSlice-4           33495             34393 ns/op            1050 B/op          2 allocs/op
BenchmarkSlice-4           33523             35001 ns/op            1050 B/op          2 allocs/op
BenchmarkSlice-4           35238             34105 ns/op            1050 B/op          2 allocs/op
BenchmarkSlice-4           33674             33890 ns/op            1050 B/op          2 allocs/op
PASS
ok      maptests        13.751s
```

Avg performance:
- `BenchmarkMap-4`: **1832/s**
- `BenchmarkSlice-4`: **29130/s**

Map is **1503%** (15 times) slower than slice **1503%**

Comparison with [benchstat](https://pkg.go.dev/golang.org/x/perf/cmd/benchstat) on more runs:

```shell
$ benchstat ./slice.txt ./map.txt
goos: linux
goarch: amd64
pkg: maptests
cpu: Intel(R) Core(TM) i7-1065G7 CPU @ 1.30GHz
           │ ./slice.txt │               ./map.txt                │
           │   sec/op    │    sec/op     vs base                  │
MapSlice-4   33.96µ ± 3%   616.04µ ± 4%  +1713.93% (p=0.000 n=10)

           │ ./slice.txt  │               ./map.txt               │
           │     B/op     │     B/op      vs base                 │
MapSlice-4   1.025Ki ± 0%   5.369Ki ± 0%  +423.57% (p=0.000 n=10)

           │ ./slice.txt │              ./map.txt              │
           │  allocs/op  │ allocs/op   vs base                 │
MapSlice-4    2.000 ± 0%   4.000 ± 0%  +100.00% (p=0.000 n=10)
```



## Perl

```sh
$ perl -v
This is perl 5, version 32, subversion 1 (v5.32.1) built for x86_64-linux-gnu-thread-multi

$ perl ./main.pl
                   Rate  count_chars_hash count_chars_array
count_chars_hash  452/s                --              -45%
count_chars_array 826/s               83%                --
$ perl ./main.pl
                   Rate  count_chars_hash count_chars_array
count_chars_hash  461/s                --              -44%
count_chars_array 826/s               79%                --
$ perl ./main.pl
                   Rate  count_chars_hash count_chars_array
count_chars_hash  450/s                --              -43%
count_chars_array 794/s               76%                --
$ perl ./main.pl
                   Rate  count_chars_hash count_chars_array
count_chars_hash  422/s                --              -50%
count_chars_array 840/s               99%                --
$ perl ./main.pl
                   Rate  count_chars_hash count_chars_array
count_chars_hash  424/s                --              -45%
count_chars_array 769/s               82%                --
```

Avg performance:
- `count_chars_hash`: **442/s**
- `count_chars_array`: **811/s**

Map is **83%** slower than hash



## Python

```shell
$ python3 -V
Python 3.9.2

$ python3 ./main.py
count_characters_dict
4.02568349899957
count_characters_array
2.8168203800014453

$ python3 ./main.py
count_characters_dict
3.962258641993685
count_characters_array
2.828529840000556

$ python3 ./main.py
count_characters_dict
3.8322311189986067
count_characters_array
2.82430825499614

$ python3 ./main.py
count_characters_dict
3.8427813529997366
count_characters_array
2.9834478900011163

$ python3 ./main.py
count_characters_dict
3.9299439580063336
count_characters_array
2.837937961005082
```

Avg performance:
- `count_characters_dict`: **255/s**
- `count_characters_array`: **350/s**

Dictionary is **37%** slower than list
