.PHONY: lint
# lint
lint:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	golangci-lint run