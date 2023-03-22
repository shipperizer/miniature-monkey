.PHONY=test build

GO111MODULE?=on
CGO_ENABLED?=0
GOOS?=linux
GO?=go
GO_BIN?=/go/bin/app
MAIN_DIR?=/var/app
APP_NAME?=web


.EXPORT_ALL_VARIABLES:

build:
	$(GO) build -a -installsuffix nocgo -buildvcs=true -o $(GO_BIN) example/$(APP_NAME)/cmd
	# $(MAKE) -C example/$(APP_NAME)/cmd build

test: mocks vet
	$(GO) test ./... -cover -coverprofile coverage.source.out
	# this will be cached, just needed to the test.json
	$(GO) test ./... -cover -coverprofile coverage.source.out -json > test.source.json
	cat coverage.source.out | grep -v "mock_*" | tee coverage.out
	cat test.source.json | grep -v "mock_*" | tee test.json

mocks: vendor
	$(GO)  install github.com/golang/mock/mockgen@v1.6.0
	# generate gomocks
	$(GO) generate ./...

vet:
	$(GO) vet ./...

vendor:
	$(GO) mod vendor
