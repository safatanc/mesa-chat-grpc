generate:
	protoc --go_out=pb/user --go_opt=paths=source_relative --go-grpc_out=pb/user --go-grpc_opt=paths=source_relative proto/user.proto
	protoc --go_out=pb/chat --go_opt=paths=source_relative --go-grpc_out=pb/chat --go-grpc_opt=paths=source_relative proto/chat.proto

run:
	go run server.go