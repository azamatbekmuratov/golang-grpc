# golang-grpc
Golang  gRPC showcase

#### To generate protobuf files use the command:

protoc --proto_path=proto proto/*.proto --go_out=. --go-grpc_out=. 

***

#### To install grpcul use the command:

protoc --proto_path=proto proto/*.proto --go_out=. --go-grpc_out=.

***

#### To call the rpc method use the command:

grpcurl -plaintext localhost:8080 Inventory/GetBookList

***

#### To list available RPC methods use the command through reflection:

grpcurl -import-path proto -proto bookshop.proto list

***