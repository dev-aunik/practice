package main

import (
	"log"
	"practice/cmd/api"
)

func main() {
	server := api.NewAPIServer(":80", nil)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
