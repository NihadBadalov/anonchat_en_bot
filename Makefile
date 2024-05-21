build:
	@go build -o bin/anonchat_en_bot ./cmd/anonchat_en_bot

run: build
	@./bin/anonchat_en_bot

test:
	@go test -v ./... -v
