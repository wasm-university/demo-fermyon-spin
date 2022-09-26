package main

import (
	hf "github.com/bots-garden/capsule/capsulemodule/hostfunctions"
)

func main() {
    hf.SetHandle(Handle)
}

func Handle(params []string) (string, error) {

  arguments := `{"x":10,"y":32}`

  sum, err := hf.NatsConnectRequest("localhost:4222", "maths", arguments, 1)

  return arguments + " sum:" + sum, err
}
