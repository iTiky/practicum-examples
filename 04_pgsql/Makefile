lint:
	@echo "Running linter checks"
	golangci-lint run

test:
	@echo "Running UNIT tests"
	@go clean -testcache
	go test -cover -race -short ./... | { grep -v 'no test files'; true; }

build:
	@echo "Building the app to the .build dir"
	go build -o .build/psql-example ./cmd/psql-example/main.go
