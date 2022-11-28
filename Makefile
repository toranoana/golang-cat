BIN_NAME := golang-cat
BIN_DIR := ./bin

GOBIN ?= $(shell go env GOPATH)/bin

.PHONY: all
all: build

.PHONY: build
build:
	mkdir -p $(BIN_DIR)
	go build -o $(BIN_DIR)/$(BIN_NAME) main.go


.PHONY: build
x-build:
	@echo 'クロスコンパイルをgoxzでやる'

git-release:
	@echo 'ghrを使ってreleaseを作る'

$(GOBIN)/goxz:
	@echo 'goxzのインストール'

$(GOBIN)/ghr:
	@echo 'ghrのインストール'
