package main

import (
	"strconv"

	"github.com/bots-garden/capsule/capsulemodule/flatjson"
	hf "github.com/bots-garden/capsule/capsulemodule/hostfunctions"
)

func main() {
	hf.OnNatsMessage(Handle)
}

func Handle(params []string) {
	natsMessage := params[0]

  hf.Log("🖐 on subject: " + hf.NatsGetSubject() + ", 🎉 message: " + natsMessage)

  arguments := flatjson.StrToMap(natsMessage)

  x := arguments["x"].(int)
  y := arguments["y"].(int)
  sum := x + y

  hf.NatsReply(strconv.Itoa(sum), 10)

}

//export OnLoad
func OnLoad() {
	hf.Log("🙂 Hello from the NATS subscriber")
	hf.Log("👂Listening on: " + hf.NatsGetSubject())
	hf.Log("👋 NATS server: " + hf.NatsGetServer())

}

//export OnExit
func OnExit() {
	hf.Log("👋🤗 Bye! Have a nice day")
}
