#bin/sh
$GOPATH/bin/mockgen -source=dom/dom.go -destination=mock/dom.go -package=mock
$GOPATH/bin/mockgen -source=node.go -destination=mock/node.go -package=mock
GOOS=js GOARCH=wasm go test ./...