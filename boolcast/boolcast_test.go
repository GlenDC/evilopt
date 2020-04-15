package boolcast

import (
	"testing"
)

var benchmarkInputs = []bool{true, false}

func benchmarkCastFunc(b *testing.B, f func(bool) byte) {
	for n := 0; n < b.N; n++ {
		f(benchmarkInputs[n%2])
	}
}

func BenchmarkCastUnsafe(b *testing.B) {
	benchmarkCastFunc(b, CastUnsafe)
}

func BenchmarkCastConditional(b *testing.B) {
	benchmarkCastFunc(b, CastConditional)
}

func BenchmarkCastMapped(b *testing.B) {
	benchmarkCastFunc(b, CastMapped)
}

func BenchmarkCastStandard(b *testing.B) {
	benchmarkCastFunc(b, CastStandard)
}

func BenchmarkCastFoundationDB(b *testing.B) {
	benchmarkCastFunc(b, CastFoundationDB)
}
