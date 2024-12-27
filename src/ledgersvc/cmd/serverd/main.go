package main

import (
	"log"

	"github.com/kneadCODE/fursave/src/golib/config"
)

func main() {
	log.Println("Welcome to Fursave")

	_, err := config.Init()
	if err != nil {
		log.Fatalf("Failed to initialize App: %v", err)
	}
}
