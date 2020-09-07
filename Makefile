all: generate typecheck test server

generate:
	go generate -x ./...

typecheck:
	go build ./...

server: generate
	go build ./app/cmd/server

playground: generate
	go run ./app/cmd/playground

unit-test: generate
	go test -v -race ./app/...

e2e: generate
	go test -v -race -tags e2e ./app/e2e/...

test: unit-test e2e
