BINDIR	:= $(CURDIR)/bin
SRC	:= $(shell find . -type f -name '*.go' -print) go.mod go.sum
GOFLAGS	:= -trimpath

.PHONY: all build clean lint test

all: build

# TODO Use this if cmd/main.go used
build: $(SRC)
	go build $(GOFLAGS) -o $(BINDIR)/example ./cmd

# TODO Use this if cmd/<names> used
#build: $(SRC)
#	go build $(GOFLAGS) -o $(BINDIR)/ ./cmd/...

clean:
	@rm -rf '$(BINDIR)'

lint:
	golangci-lint run

test:
	go test ./...
