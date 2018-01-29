package test

import (
	"testing"
	"fmt"
)

var MinLevel = 20

func Trace(kv ...string) {
	if 10 < MinLevel {
		return
	}
	fmt.Println(append(([]string)(nil), kv...))
}

func Benchmark_arguments_passing(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		Trace("hello", "world", "hey", "there")
	}
}
