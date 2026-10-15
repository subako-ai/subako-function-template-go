# Subako function package template (Go)

A starting point for a [Subako](https://subako.ai) function package written in Go.

```sh
go run golang.org/x/tools/cmd/gonew@latest github.com/subako-ai/subako-function-template-go example.com/hello
cd hello
```

`gonew` copies the template and renames its module. Give the package its own name and description in `subako-package.json`, which Subako reads when the package is uploaded.

## Write tools

Each tool is a Go function added to an [official MCP Go SDK](https://github.com/modelcontextprotocol/go-sdk) server with `mcp.AddTool`:

- the tool's `Name` is the name the agent calls it by;
- its `Description` is what the model reads;
- the function's input and output structs become the tool's schemas, with each field's `jsonschema` tag as its description.

An error the function returns comes back to the model as a failed call it can read. `subakofn.Serve` serves the tools as an MCP server over stdio (Subako's function ABI 1). Inside a call, the session's Volume is at `/volume/tmp` and `/volume/data`, and the package's own `assets/` directory is at `/package`. `net/http` sends requests through Subako, which applies the function's network setting and credentials, and `subakofn.ImageResult` shows the model an image from the Volume. See [subako-function-sdk-go](https://github.com/subako-ai/subako-function-sdk-go) for the rest.

## Test

`go test ./...` runs your tools as plain Go.

To try the module itself with [wasmtime](https://wasmtime.dev), which has no `subako_http` import, let `tests/subako_http_stub.wat` stand in for it:

```sh
GOOS=wasip1 GOARCH=wasm go build -o dist/module.wasm .
printf '%s\n' '{"jsonrpc":"2.0","id":1,"method":"tools/call","params":{"_meta":{"io.modelcontextprotocol/protocolVersion":"2026-07-28","io.modelcontextprotocol/clientCapabilities":{}},"name":"greet","arguments":{"name":"Ada"}}}' |
    wasmtime run --preload subako_http=tests/subako_http_stub.wat dist/module.wasm
```

## Package and upload

```sh
bash package.sh                                # dist/hello.tar.gz
subako function create dist/hello.tar.gz       # first version
subako function push <function_id> dist/hello.tar.gz   # later versions
```

Then grant the package to an agent under **Function packages** in the agent version editor.

`package.sh` needs Go 1.25 or later. It runs through `bash` because `gonew` does not keep files executable.

## License

Licensed under either of [Apache License, Version 2.0](LICENSE-APACHE) or [MIT license](LICENSE-MIT) at your option.
