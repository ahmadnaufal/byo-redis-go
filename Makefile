.PHONY: build run testclient

build:
	go build -o bin/server .

run:
	go run .

testclient:
	go run ./testclient/main.go