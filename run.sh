#!/bin/sh
cd ~/goPrograms/wails-to-wasm/cmd/wasm
GOOS=js GOARCH=wasm go build -o  ../../frontend/json.wasm
cd ~/goPrograms/wails-to-wasm/cmd/server
go run main.go

