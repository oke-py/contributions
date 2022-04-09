GO := GO111MODULE=on go
GOBIN := $(shell go env GOPATH)/bin

all: fix vet fmt test build tidy

build:
	$(GO) build -o bin/contribution cmd/main.go

.PHONY: fix
fix:
	$(GO) fix ./pkg/... ./cmd/...

.PHONY: fmt
fmt:
	$(GO) fmt ./pkg/... ./cmd/...

test:
	$(GO) test -v -covermode=count -coverprofile=coverage.out ./...

tidy:
	$(GO) mod tidy

.PHONY: vet
vet:
	$(GO) vet ./pkg/... ./cmd/...
