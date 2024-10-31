package main

import (
	"MyStudy/server"
	"log"
	"net/http"
)

func main() {

	hostName := ":5000"
	server.Server()
	log.Fatal(http.ListenAndServe(hostName, nil))
}
