package main

import (
	"log"

	"github.com/FelipeAz/golibcontrol/build/server/account/server"
)

func main() {
	err := server.Start()
	if err != nil {
		log.Println(err.Error())
	}
}
