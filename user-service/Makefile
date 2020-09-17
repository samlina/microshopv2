
GOPATH:=$(shell go env GOPATH)
MODIFY=Mgithub.com/micro/go-micro/api/proto/api.proto=github.com/micro/go-micro/v2/api/proto

.PHONY: proto
proto:
    
	protoc --proto_path=. --micro_out=${MODIFY}:. --go_out=${MODIFY}:. proto/user/user.proto
    

.PHONY: build
build: proto

	go build -o user-service *.go

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: docker
docker:
	docker build . -t user-service:latest


run:
	 docker run -d -p 9091:9091 -e MICRO_SERVER_ADDRESS=:9091 -e MICRO_REGISTRY=mdns user-service