#!/bin/zsh
protoc \
        --proto_path=api/proto/customer \
        --proto_path=third_party \
        --go_out=pkg/api \
        --go-grpc_out=pkg/api \
        --govalidators_out=pkg/api \
        --grpc-gateway_out=logtostderr=true:pkg/api \
        customer.proto

protoc \
      --proto_path=api/proto/customer \
      --proto_path=third_party \
      --swagger_out=logtostderr=true:api/swagger \
      customer.proto

