package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	port := os.Getenv("port")

	log.Printf("Got port *:%s\n", port)
	fmt.Println("Hello world")
}
