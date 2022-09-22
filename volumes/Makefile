pb:
	protoc --go_out=. --go-grpc_out=require_unimplemented_servers=false:. ./proto/api.proto

build:
	go build cmd/server/main.go