PROTOC_VERSION=3.9.1
PROTOC_ZIP=protoc-${PROTOC_VERSION}-linux-x86_64.zip
PROTOC_PATH= \
    --proto_path=. \
	--proto_path=${GOPATH}/src/ \
	--proto_path=${GOPATH}/src/github.com/amsokol/protoc-gen-gotag/ \
	--proto_path=${GOPATH}/src/github.com/protocolbuffers/protobuf/src/ \
	--proto_path=${GOPATH}/src/github.com/golang/protobuf/protoc-gen-go/ \
	--proto_path=${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis/ \
	--proto_path=${GOPATH}/src/github.com/grpc-ecosystem/grpc-gateway/

PROTOC_PLUGINS=grpc

all:

.PHONY: proto auth up down deps redoc

proto:
	protoc $(PROTOC_PATH) \
		--grpc-gateway_out=logtostderr=true:. \
		--go_out=plugins=$(PROTOC_PLUGINS):. \
		./core/test/*.proto
	# protoc $(PROTOC_PATH) --gotag_out=xxx="bson+\"-\"",output_path=./core/test/pb:. ./core/test/*.proto

up:
	docker-compose up -d --build

down:
	docker-compose down

build:
	go build -o serialization github.com/wcrbrm/grpc-envoy-example/cmd/serialization

deps:
	curl -OL https://github.com/google/protobuf/releases/download/v${PROTOC_VERSION}/${PROTOC_ZIP}
	sudo unzip -o ${PROTOC_ZIP} -d /usr/local bin/protoc
	sudo unzip -o ${PROTOC_ZIP} -d /usr/local include/*
	rm -f ${PROTOC_ZIP}
	go get -u github.com/golang/protobuf/protoc-gen-go
	go get -u github.com/amsokol/protoc-gen-gotag
	go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
	go get -u google.golang.org/grpc
	go get -u github.com/amsokol/mongo-go-driver-protobuf

