.PHONY: help
help:
	@echo "API server on port 8080"

run:
	go run ./cmd/main/app.go