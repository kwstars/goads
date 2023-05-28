.PHONY: lint
# lint
lint:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	golangci-lint run

.PHONY: lint-fix
lint-fix:
	golangci-lint run --concurrency=2 --fix

.PHONY: test
test:
	go test -v -count=1 -covermode=atomic -coverprofile=.testCoverage.txt -timeout=2m ./...
	go test -v ./... -race


.PHONY: test-v
test-v:
	go test -count=1 -race -covermode=atomic -coverprofile=.testCoverage.txt -timeout=2m -v ./...

.PHONY: cover-view
cover-view:
	go tool cover -func .testCoverage.txt
	go tool cover -html .testCoverage.txt

.PHONY: build
build:
	go build ./...