all: generate typecheck test server

clear:
	rm -rf build

generate:
	go generate -x ./...

typecheck:
	go build ./...

build-dir:
	mkdir -p build

server: generate build-dir
	go build -o build/server ./app/cmd/server

playground: generate
	go run ./app/cmd/playground

unit-test: generate
	go test -v -race ./app/...

e2e: generate
	go test -v -race -tags e2e ./app/e2e/...

test: unit-test e2e
