all: generate typecheck test server

generate:
	go generate -x ./...

typecheck:
	go build ./...

server: generate
	go build ./app/cmd/server

playground: generate
	go run ./app/cmd/playground

test: generate
	go test -v -race ./app/...
