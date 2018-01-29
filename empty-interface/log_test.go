package test

import (
	"testing"
	"fmt"
	"unsafe"
)

var MinLevel = 20

func Trace(kv ...interface{}) {
	if 10 < MinLevel {
		return
	}
	ptr := unsafe.Pointer(&kv)
	ptrAsValue := uintptr(ptr)
	args := castEmptyInterfaces(ptrAsValue)
	fmt.Println(args)
}

func castEmptyInterfaces(ptr uintptr) []interface{} {
	return *(*[]interface{})(unsafe.Pointer(ptr))
}

var k1 = "k1"
var v1 = 100
var k2 = "k2"
var v2 = 10.24

func Benchmark_empty_interface(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		Trace(k1, v1, k2, v2)
	}
}
