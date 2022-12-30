# Belajar GRPC


## Prerequisites

- Install protoc compiler & Go plugging https://grpc.io/docs/languages/go/quickstart/
- Add path ```export PATH="$PATH:$(go env GOPATH)/bin"```

## FYI

1. Need define ```option go_package``` in your proto file, and use ```./```
2. For generate ```*.pb.go``` & ```_grpc.pb.go``` file, run this command

    ```protoc --proto_path=model/proto model/proto/*.proto --go-grpc_out=. --go_out=.```

    that will generape file in model folder