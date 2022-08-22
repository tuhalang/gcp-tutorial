test:
	go test -v -cover ./...
build:
	go build -o main
run:
	go run main.go
.PHONY: test build run