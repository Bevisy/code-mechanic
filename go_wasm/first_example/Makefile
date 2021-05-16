.PHONY: run test server build clean

run:
	GOOS=js GOARCH=wasm go run -exec="$$(go env GOROOT)/misc/wasm/go_js_wasm_exec" ./example

test:
	GOOS=js GOARCH=wasm go test -exec="$$(go env GOROOT)/misc/wasm/go_js_wasm_exec" ./example

server: build
	echo "https://127.0.0.1:8080"
	go run localDirServer.go

build:
	GOOS=js GOARCH=wasm go build -o main.wasm example/main.go

clean:
	rm -f main.wasm