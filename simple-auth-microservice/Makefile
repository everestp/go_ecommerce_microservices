run:
	go run ./cmd/auth/main.go

build:
	go build -o simple-auth-microservice ./cmd/auth/main.go

proto:
	protoc --go_out=./proto --go-grpc_out=./proto --proto_path=./proto ./proto/*.proto