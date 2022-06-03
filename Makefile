generate-server:
	protoc --proto_path=src/protofiles --go_out=src/server --go-grpc_out=src/server $(protofile)

generate-client:
	protoc --proto_path=src/protofiles --go_out=src/client --go-grpc_out=src/client $(protofile)

generate-grpc: generate-server generate-client

start-evans:
	evans -r repl -p $(port)