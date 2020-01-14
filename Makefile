all: deps prepare-test test 
deps:
	GOOS=js GOARCH=wasm go get ./...
prepare-test:
	$(GOPATH)/bin/mockgen -source=dom/dom.go -destination=mock/dom.go -package=mock
	$(GOPATH)/bin/mockgen -source=vnode.go -destination=mock/node.go -package=mock
test:
	GOOS=js GOARCH=wasm go test ./... -exec="$(shell go env GOROOT)/misc/wasm/go_js_wasm_exec"
test-cover:
	GOOS=js GOARCH=wasm go test ./... -exec="$(shell go env GOROOT)/misc/wasm/go_js_wasm_exec" -coverprofile=cover.out
	go tool cover -html=cover.out