# wails-to-wasm

This repository is designed to create files that can be used in creating a web page based on source code (go, html, css) from the [Hello World tutorial](https://wails.io/docs/tutorials/helloworld). This approach uses webAssembly(wasm) technology. It is based on an approach outlined in [this tutorial](https://golangbot.com/webassembly-using-go/) and [it's subsequent tutorial](https://golangbot.com/go-webassembly-dom-access/).

## Reusing Existing Code

The go code was taken from app.go.  The body of the html was taken from frontend/src/main.js.  The css was taken from frontend/src/app.css and frontend/src/style.css.  The logo-universal.png was taken from frontend/src/assets/images. Some minor modifications were necessary. I tried to leave the original code as a comment when the modifications were made.

## Testing

Run the commands found in run.sh.

Go to a browser and enter ```localhost:9090``` in the address window.

## Deployment

A example of a deployment method can be found [here](https://github.com/tataDan/wails-to-wasm-app/).

The actual app can be found [here](https://tatadan.github.io/wails-to-wasm-app/).
