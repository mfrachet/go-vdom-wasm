all: test build
build:
	GOOS=js GOARCH=wasm go get ${gobuild_args} ./...
test: 
	GOOS=js GOARCH=wasm go test ./...