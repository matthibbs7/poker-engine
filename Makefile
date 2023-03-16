build:
	@go build -o bin/poker-engine

run: build
	@./bin/poker-engine

test:
	go test -v ./...