package main

func main() {
	println("hello world! tinygo wasm")
}

/**
*	magic!! Tinygo use //export comment header
* to indicate which function will be exported
 */

//export add
func add(x, y int) int {
	return x + y
}

//export substract
func substract(x, y int) int {
	return x - y
}
