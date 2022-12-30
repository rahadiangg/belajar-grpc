generate-model:
	protoc --proto_path=model/proto model/proto/*.proto --go-grpc_out=. --go_out=.
	@echo "Model created"

run-server:
	go run main.go

run-client:
	go run client.go