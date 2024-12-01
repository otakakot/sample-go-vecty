# sample-go-vecty

```shell
cp $(go env GOROOT)/misc/wasm/wasm_exec.js .
```

```shell
cp $(go env GOROOT)/misc/wasm/wasm_exec.html index.html
```

```shell
GOOS=js GOARCH=wasm go build -ldflags "-s -w" -trimpath -o main.wasm main.go
```
