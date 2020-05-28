lint:
	golangci-lint run

import:
		goimports -w -l .

format:
		go fmt ./...

test:
		go test ./... -v

all: format lint test