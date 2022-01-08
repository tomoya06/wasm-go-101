package main

import (
	"fmt"
	"syscall/js"
)

func add(this js.Value, args []js.Value) interface{} {
	a, b := args[0].Int(), args[1].Int()
	return a + b
}

func substract(this js.Value, args []js.Value) interface{} {
	a, b := args[0].Int(), args[1].Int()
	return a - b
}

func main() {
	fmt.Println("hello wasm world")

	js.Global().Set("add", js.FuncOf(add))
	js.Global().Set("substract", js.FuncOf(substract))

	select {}
}
