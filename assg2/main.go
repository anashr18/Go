package main

import (
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	file, err := os.OpenFile("input.txt", os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	defer file.Close()

	r := strings.NewReader("hhjjjh")
	_, er := io.Copy(file, r)
	if er != nil {
		log.Fatalf("Failed to copy content to the file: %v", er)
	}
}
