GOOS ?= $(shell go env GOOS)
GOARCH ?= $(shell go env GOARCH)


.PHONY: genall
genall:
	./scripts/genproto.sh
	./scripts/genswagger.sh

.PHONY: genproto
genproto:
	./scripts/genproto.sh

.PHONY: genswagger
genswagger:
	./scripts/genswagger.sh


.PHONY: test
test:
	go test ./cmd/...

.PHONY: clean
clean:
	./scripts/clean.sh

.PHONY: install
install:
	go env -w GOPROXY=https://goproxy.cn,direct && \
	go get \
	github.com/gogo/protobuf/protoc-gen-gogo \
	github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway \
	github.com/grpc-ecosystem/grpc-gateway/protoc-gen-swagger \
	github.com/mwitkow/go-proto-validators/protoc-gen-govalidators \
	github.com/rakyll/statik