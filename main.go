package main

import (
	"gin-demo/adapter"
	"log"
)

func main() {
	if err := adapter.Init(); err != nil {
		log.Fatal("Failed to initialize server: ", err)
	}
}
