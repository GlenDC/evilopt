# Boolcast

What doesn't require any work in C, does require work in a language such as Go.

## Golang Results

Sorted from best forming to least performing (average of 5 runs)

1. `CastUnsafe`: 1.928 ns/op
2. `CastConditional`: 2.154 ns/op
3. `CastMapped`: 19.58 ns/op
4. `CastStandard`: 25.82 ns/op
5. `CastFoundationDB`: 42.22 ns/op

Making this ugly little one the winner in land of the gophers:

```go
import "unsafe"

func CastUnsafe(b bool) byte {
	return *(*byte)(unsafe.Pointer(&b)) & 1
}
```

The conditional is pretty much equivalent in terms of performance though,
so going for that obvious solution isn't too shabby either.

Runs:
```
for i in {1..5}; do
    /usr/local/bin/go test -benchmem -run='^$' github.com/evilopt/boolcast -bench '^(BenchmarkCastUnsafe|BenchmarkCastConditional|BenchmarkCastMapped|BenchmarkCastStandard|BenchmarkCastFoundationDB)$';
done
goos: darwin
goarch: amd64
pkg: github.com/evilopt/boolcast
BenchmarkCastUnsafe-12                  621978982                1.94 ns/op            0 B/op          0 allocs/op
BenchmarkCastConditional-12             552589515                2.18 ns/op            0 B/op          0 allocs/op
BenchmarkCastMapped-12                  61054688                19.9 ns/op             0 B/op          0 allocs/op
BenchmarkCastStandard-12                42221215                25.4 ns/op             1 B/op          1 allocs/op
BenchmarkCastFoundationDB-12            26800370                41.0 ns/op            64 B/op          1 allocs/op
PASS
ok      github.com/evilopt/boolcast     6.480s
goos: darwin
goarch: amd64
pkg: github.com/evilopt/boolcast
BenchmarkCastUnsafe-12                  618867904                1.90 ns/op            0 B/op          0 allocs/op
BenchmarkCastConditional-12             562479447                2.14 ns/op            0 B/op          0 allocs/op
BenchmarkCastMapped-12                  58785108                19.4 ns/op             0 B/op          0 allocs/op
BenchmarkCastStandard-12                44318568                25.3 ns/op             1 B/op          1 allocs/op
BenchmarkCastFoundationDB-12            25509654                41.2 ns/op            64 B/op          1 allocs/op
PASS
ok      github.com/evilopt/boolcast     6.372s
goos: darwin
goarch: amd64
pkg: github.com/evilopt/boolcast
BenchmarkCastUnsafe-12                  617770444                1.92 ns/op            0 B/op          0 allocs/op
BenchmarkCastConditional-12             548001534                2.11 ns/op            0 B/op          0 allocs/op
BenchmarkCastMapped-12                  58385846                19.5 ns/op             0 B/op          0 allocs/op
BenchmarkCastStandard-12                41343774                25.4 ns/op             1 B/op          1 allocs/op
BenchmarkCastFoundationDB-12            27889286                42.8 ns/op            64 B/op          1 allocs/op
PASS
ok      github.com/evilopt/boolcast     6.396s
goos: darwin
goarch: amd64
pkg: github.com/evilopt/boolcast
BenchmarkCastUnsafe-12                  575774870                1.94 ns/op            0 B/op          0 allocs/op
BenchmarkCastConditional-12             556910844                2.15 ns/op            0 B/op          0 allocs/op
BenchmarkCastMapped-12                  59873050                19.7 ns/op             0 B/op          0 allocs/op
BenchmarkCastStandard-12                43311624                27.0 ns/op             1 B/op          1 allocs/op
BenchmarkCastFoundationDB-12            25769242                43.3 ns/op            64 B/op          1 allocs/op
PASS
ok      github.com/evilopt/boolcast     7.348s
goos: darwin
goarch: amd64
pkg: github.com/evilopt/boolcast
BenchmarkCastUnsafe-12                  591007356                1.94 ns/op            0 B/op          0 allocs/op
BenchmarkCastConditional-12             554012031                2.19 ns/op            0 B/op          0 allocs/op
BenchmarkCastMapped-12                  60454591                19.4 ns/op             0 B/op          0 allocs/op
BenchmarkCastStandard-12                43530202                26.0 ns/op             1 B/op          1 allocs/op
BenchmarkCastFoundationDB-12            27280374                42.8 ns/op            64 B/op          1 allocs/op
PASS
ok      github.com/evilopt/boolcast     6.536s
```
