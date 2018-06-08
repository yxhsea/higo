build:
    protoc --proto_path=. --micro_out=. --go_out=. user.proto
    GOOS=linux GOARCH=amd64 go build
    docker build -t user-service .

run:
	docker run -p 50051:50051 \
	 -e MICRO_SERVER_ADDRESS=:50051 \
	 -e MICRO_REGISTRY=mdns \
	 user-service