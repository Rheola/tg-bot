.PHONY: run
run:
	go run cmd/bot/main.go

.PHONY: deps
deps:
	@go mod download
	@go mod vendor
	@go mod tidy