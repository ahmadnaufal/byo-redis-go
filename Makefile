.PHONY: build run testclient

build:
	go build -o ./bin/testserver ./cmd/testserver/main.go
	go build -o ./bin/testclient ./cmd/testclient/main.go

server:
	go run ./cmd/testserver/main.go

testclient:
	go run ./cmd/testclient/main.go