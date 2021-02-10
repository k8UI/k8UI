lint:
	@echo "Running golangci-lint..."
	@golangci-lint run

test:
	@echo "Running tests..."
	@go test ./internal/... -cover -count=1