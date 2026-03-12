build:
	@go build -o bin/kvstore

run: build
	@./bin/kvstore	