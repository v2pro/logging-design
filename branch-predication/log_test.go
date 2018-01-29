package test

import "testing"

func Benchmark_trace_branch_predication(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		Trace("hello")
	}
}