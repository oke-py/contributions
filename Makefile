GO := GO111MODULE=on go
GOBIN := $(shell go env GOPATH)/bin

all: fix vet fmt lint sec test build tidy

build:
	$(GO) build -o bin/contribution cmd/main.go

.PHONY: fix
fix:
	$(GO) fix ./pkg/... ./cmd/...

.PHONY: fmt
fmt:
	$(GO) fmt ./pkg/... ./cmd/...

lint:
	(which $(GOBIN)/golangci-lint || go get github.com/golangci/golangci-lint/cmd/golangci-lint@v1.35.2)
	$(GOBIN)/golangci-lint run ./...

sec:
	(which $(GOBIN)/gosec || go get github.com/securego/gosec/cmd/gosec)
	$(GOBIN)/gosec ./pkg/... ./cmd/...

test:
	$(GO) test -v -covermode=count -coverprofile=coverage.out ./...

tidy:
	$(GO) mod tidy

.PHONY: vet
vet:
	$(GO) vet ./pkg/... ./cmd/...
