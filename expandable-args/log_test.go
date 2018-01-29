package test

import (
	"testing"
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
	for i := 0; i < len(args); i += 2 {
		key := args[i].(string)
		value := args[i+1].(func() interface{})()
		blackHole(key)
		blackHole(value)
	}
}

func blackHole(interface{}) {
}

func castEmptyInterfaces(ptr uintptr) []interface{} {
	return *(*[]interface{})(unsafe.Pointer(ptr))
}

var v1 = 1024 * 1024
var v2 = 4096 * 4096

func Benchmark_expandable_args(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		Trace("k1", func() interface{} {
			return make([]byte, v1)
		}, "v2", func() interface{} {
			return make([]byte, v2)
		})
	}
}