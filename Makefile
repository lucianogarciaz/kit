lint:
	golangci-lint run
	go mod tidy -v && git --no-pager diff go.mod go.sum

lint-fix:
	golangci-lint run --fix

test:
	go test -race ./...

tool-moq:
	scripts/goget.sh github.com/matryer/moq@v0.2.1

mocks: ## Create mocks
	go mod vendor
	# cqs
	moq -out cqs/zmock_cqs_test.go -pkg cqs_test cqs QueryHandler CommandHandler Command Event EventHandler
	moq -out obs/zmock_obs_test.go -pkg obs_test obs Marshaler Observer
	moq -out obs/zmock_cqs_test.go -pkg obs_test cqs Command CommandHandler Query QueryHandler
	rm -rf ./vendor