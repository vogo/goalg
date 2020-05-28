lint:
	golangci-lint run

import:
		goimports -w -l .

format:
		go fmt ./...

test:
		go test ./... -v

ci: format test

all: import format lint test