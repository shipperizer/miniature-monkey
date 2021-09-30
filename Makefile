.PHONY=test build

GOPRIVATE?=github.com/sbs-discovery-sweden/*
GO111MODULE?=on
CGO_ENABLED?=0
GOOS?=linux
GO?=go
GO_BIN?=/go/bin/app
MAIN_DIR?=/var/app
APP_NAME?=test


.EXPORT_ALL_VARIABLES:

build:
	$(MAKE) -C cmd/$(APP_NAME) build

test: mocks
	$(GO) test -v ./... -cover -coverprofile=coverage.out -covermode=atomic
	cat coverage.out | grep -v "mock_" > coverage.out

mocks:
	$(GO) get github.com/golang/mock/gomock
	$(GO) get github.com/golang/mock/mockgen
	# generate gomocks
	$(GO) generate ./...
