## Test call http.Request from wasm in javascript.

* build wasm file: `GOOS=js GOARCH=wasm go build  -o main.wasm`
* Move to webserver assets: `mv main.wasm ./html/main.wasm`
* Run http server: `go run ./cmd/`