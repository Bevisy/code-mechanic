package main

import (
	"syscall/js"
)

func main() {
	console_log := js.Global().Get("console").Get("log")
	console_log.Invoke("Hello wasm!")

	js.Global().Call("eval", `
        console.log("hello, wasm!");
    `)
}
