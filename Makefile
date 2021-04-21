VERSION = "0.0.3"
change-version:
	@echo $(VERSION)>VERSION
	@echo "package constant\n\n//Version constant of the service\nconst Version = \"$(VERSION)\"">pkg/constant/version.go

update-module:
	go get -v google.golang.org/grpc
	go get -v github.com/rs/zerolog
	go get -v github.com/caarlos0/env/v6
	go get -v google.golang.org/protobuf
	go get -v github.com/mwitkow/go-proto-validators
	go get -v github.com/satori/go.uuid
	go get -v github.com/grpc-ecosystem/grpc-gateway
	go get -v github.com/go-pg/pg/v10

build:
	GOPRIVATE=repo GIT_TERMINAL_PROMPT=1 CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -o bin/grpc_server cmd/server/server.go

simple-build:
	go build -v -o bin/grpc_server cmd/server/server.go

clean:
	@rm -rf bin/grpc_server
	@rm -rf .proto-dir

run:
	go run cmd/server/server.go


protoc:
	@protoc \
        --proto_path=api/proto/customer \
        --proto_path=third_party \
        --go_out=pkg/api \
        --go-grpc_out=pkg/api \
        --govalidators_out=pkg/api \
        --grpc-gateway_out=logtostderr=true:pkg/api \
        customer.proto