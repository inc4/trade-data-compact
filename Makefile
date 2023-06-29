BINDIR	:= bin
SRC	:= $(shell find . -type f -name '*.go' -print) go.mod go.sum
GOFLAGS	:= -trimpath

.PHONY: all build clean lint test

all: build

build: $(SRC)
	go build $(GOFLAGS) -o $(BINDIR)/compact ./cmd

clean:
	@rm -rf '$(BINDIR)'

lint:
	golangci-lint run

test:
	go test ./...
