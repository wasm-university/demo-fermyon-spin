package main

import (
	"fmt"
	"io"
	"net/http"

	spinhttp "github.com/fermyon/spin/sdk/go/http"
)

func init() {
	spinhttp.Handle(func(w http.ResponseWriter, r *http.Request) {

    var response string
    bytes, err := io.ReadAll(r.Body)
    if err != nil {
      response = err.Error()
    } else {
      response = string(bytes)
    }

    fmt.Println(string(bytes))
    // logs:
    // bat $HOME/.spin/go-hello-world/logs/go-hello-world_stdout.txt
    // bat $HOME/.spin/go-hello-world/logs/go-hello-world_stderr.txt

		w.Header().Set("Content-Type", "text/plain")
		fmt.Fprintln(w, "👋 Hello World! 🌍", response)
	})
}

func main() {
}
