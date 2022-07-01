#! /usr/bin/make

export GO111MODULE=on

all:build test build
	@echo DONE!

dokcer:
	@echo GENERATING DOCKER...
	@docker build .

lint:
	@echo LINTING CODE...

dep:
	@echo DOWONLOADING MODULES...
	@go mod download

fmt:
	@echo FORMATTING CODE...
	@go fmt

build:
	@echo GENERATING CODE...
	@go build main.go

test:
	@echo TESTING...
	@go test ./...

serve: build
	@echo SERVING...
	@./main
