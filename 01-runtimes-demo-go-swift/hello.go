package main

import (
	"fmt"
	"os"
)

//export four
func four(number float32) float32 {
  return number * 4
}

//export hello
func hello(name string) {
  fmt.Println("👋 hello", name)
}
// WebAssembly Interface Types
// https://github.com/bytecodealliance/wasmtime/blob/main/docs/wasm-rust.md

func main() {

	args := os.Args

  fmt.Println("🖐 Args:", args)

}
