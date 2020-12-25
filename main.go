package main

import (
	"log"
	"net/http"

	"github.com/memochou1993/chat/controller"
)

func main() {
	http.HandleFunc("/", controller.Handler)

	log.Fatal(http.ListenAndServe(":80", nil))
}
