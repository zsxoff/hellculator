protoc:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative protobuf/*.proto

run_client:
	go run ./client/main.go

run_server:
	go run ./server/main.go

.PHONY: protoc run_client run_server
