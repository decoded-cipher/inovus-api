build:
	@go build -o bin/inovus-api ./cmd/...

run:
	@$(HOME)/go/bin/air

tidy:
	@go mod tidy

clean:
	@echo "Cleaning build artifacts and temporary files..."
	@rm -rf bin/ tmp/
	@mkdir -p bin tmp
	@echo "Clean complete!"

.PHONY: build run tidy clean