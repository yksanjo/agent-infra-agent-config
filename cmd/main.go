package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println("AgentConfig starting...")
	if err := initialize(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("AgentConfig ready")
}

func initialize() error {
	return nil
}
