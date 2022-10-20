package main

import (
	"fmt"
	//"os"
  "tinygo.org/x/drivers/net/http"
)

func root(w http.ResponseWriter, r *http.Request) {
  w.WriteHeader(http.StatusOK)
  fmt.Fprintf(w, `<h1>👋 Hello World 🌍</h1>`, 1)
}

func main() {

	//args := os.Args

  //fmt.Println("🖐 Args:", args)

  http.HandleFunc("/", root)

  if err := http.ListenAndServe(":9090", nil); err != nil {
		fmt.Println(err.Error())
	}

}
