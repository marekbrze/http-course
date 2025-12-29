package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/charmbracelet/log"
)

func main() {
	file, err := os.Open("messages.txt")
	if err != nil {
		log.Fatal("Error when reading file")
	}
	currentLine := ""
	for {
		buffer := make([]byte, 8)
		bytes, err := file.Read(buffer)
		if bytes == 0 {
			os.Exit(0)
		}
		if err != nil {
			log.Fatal("Error when reading buffer")
		}
		parts := strings.Split(string(buffer), "\n")
		if len(parts) > 1 {
			currentLine += strings.Join(parts[:len(parts)-1], " ")
			fmt.Printf("read: %s\n", currentLine)
			currentLine = parts[len(parts)-1]
			continue
		}
		currentLine += parts[0]
	}
}
