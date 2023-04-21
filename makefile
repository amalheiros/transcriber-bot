.PHONY: build run clean test

GOCMD=go
BINARY_NAME=transcriber-bot

# Makefile
ENV := $(PWD)/.env

# Environment variables for project
include $(ENV)

# Export all variable to sub-make
export

build: 
	mkdir -p out/bin
	GO111MODULE=on $(GOCMD) build -o out/bin/$(BINARY_NAME) .

run:
	go run cmd/main.go

clean:
	rm -fr ./out

test:
	go test -race ./...