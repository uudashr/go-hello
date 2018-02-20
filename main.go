package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("Hello, %s\n", username())
}

func username() string {
	// On Linux and Mac the username is available as simply USER
	username := os.Getenv("USER")
	// check if we got an empty string and attempt to get USERNAME instead which would be available on Windows
	if username == "" {
		username = os.Getenv("USERNAME")
	}
	// Return whatever we received
	return username
}
