package main

import (
	"fmt"
	"log"
	"os"

	"github.com/suborbital/sat/sat"
)

func main() {

	argsWithoutCaller := os.Args[1:]
	wasmModulePath := argsWithoutCaller[0] // the wasm module
	wasmModuleParameters := argsWithoutCaller[1:] // the url to call

	wasmModuleConfig, _ := sat.ConfigFromRunnableArg(wasmModulePath)

	satFunction, _ := sat.New(wasmModuleConfig, nil)

	result, err := satFunction.Exec([]byte(wasmModuleParameters[0]))
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("🎉 " + string(result.Output))

}
