build:
	@go build -o bin/inovus-api ./cmd/...

run:
	@$(HOME)/go/bin/air

tidy:
	@go mod tidy