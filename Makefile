build:
	@go build -o bin/inovus-api cmd/main.go

run: build
	@./bin/inovus-api