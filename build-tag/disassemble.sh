# https://github.com/golang/go/issues/9337
go build -tags release -gcflags="-l -l -l -l -l -l -m" main.go
go tool objdump -s main.main main