BIN_NAME := golang-cat
BIN_DIR := ./bin
X_BIN_DIR := $(BIN_DIR)/goxz
VERSION := $$(make -s app-version)

GOBIN ?= $(shell go env GOPATH)/bin

.PHONY: all
all: build

.PHONY: build
build:
	mkdir -p $(BIN_DIR)
	go build -o $(BIN_DIR)/$(BIN_NAME) main.go

.PHONY: x-build
x-build: $(GOBIN)/goxz
	goxz -d $(X_BIN_DIR) -n $(BIN_NAME) .

.PHONY: upload-binary
upload-binary: $(GOBIN)/ghr
	ghr "v$(VERSION)" $(X_BIN_DIR)

.PHONY: app-version
app-version: $(GOBIN)/gobump
	@gobump show -r .

$(GOBIN)/goxz:
	@go install github.com/Songmu/goxz/cmd/goxz@latest

$(GOBIN)/ghr:
	@go install github.com/tcnksm/ghr@latest

$(GOBIN)/gobump:
	@go install github.com/x-motemen/gobump/cmd/gobump@master
