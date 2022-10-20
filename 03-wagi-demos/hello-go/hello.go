package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("content-type: text/plain;utf-8")
	fmt.Println("")

	args := os.Args
	argsWithoutCaller := os.Args[1:]

  fmt.Println("Route and QueryString:", args)
	fmt.Println(argsWithoutCaller)

  // get the body content (for POST request)
	var reader = bufio.NewReader(os.Stdin)
	message, _ := reader.ReadString('\n')

  // POST
	fmt.Println("POST:" + message)

}
