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