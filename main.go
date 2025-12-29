package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/log"
)

func main() {
	file, err := os.Open("messages.txt")
	if err != nil {
		log.Fatal("Error when reading file")
	}
	for {
		buffer := make([]byte, 8)
		bytes, err := file.Read(buffer)
		if bytes == 0 {
			os.Exit(0)
		}
		if err != nil {
			log.Fatal("Error when reading buffer")
		}
		fmt.Printf("read: %s\n", buffer)
	}
}
