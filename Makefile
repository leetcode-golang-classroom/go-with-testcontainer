.PHONY=build

build:
	@go build -o bin/main cmd/test-client.go

run: build
	@./bin/main

test-db-client:
	@go test -v -cover ./test/integration/dbclient_test.go

test-http-client:
	@go test -v -cover ./test/integration/httpclient_test.go