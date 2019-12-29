package main

import (
	"autobench-server/server"
	"log"
)

func main() {
	s, err := server.New()
	if err != nil {
		log.Fatal(err)
	}
	s.Run(":3000")
}