test: ## run the unit and integration tests
	GOOS=js GOARCH=wasm go test -exec="$(go env GOROOT)/misc/wasm/go_js_wasm_exec" -cover