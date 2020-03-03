package main

import (
	"log"
	"net/http"

	"github.com/memochou1993/chat/controller"
)

func main() {
	http.HandleFunc("/", controller.Index)
	http.HandleFunc("/api", controller.Handler)

	log.Fatal(http.ListenAndServe(":8082", nil))
}
