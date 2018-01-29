# https://github.com/golang/go/issues/9337
go test -c -gcflags="-m" log_test.go
go tool objdump -s Benchmark_empty_interface test.test