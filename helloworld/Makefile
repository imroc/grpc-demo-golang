SHELL := /bin/bash

.PHONY: proto
proto:
	protoc --go_out=. --go_opt=paths=source_relative \
	    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
	    ./greet/greet.proto

.PHONY: run_server
run_server:
	go run server/main.go

.PHONY: run_client
run_client:
	go run client/main.go