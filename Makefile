lint:
	golangci-lint run
	go mod tidy -v && git --no-pager diff go.mod go.sum

lint-fix:
	golangci-lint run --fix

test:
	go test -race ./...
