package main

import (
	"fmt"
	"os"
	"runtime"
)

func main() {
	message := make(chan string)

	go func() { message <- "ping" }()

	msg := <-message
	fmt.Println(msg)
	var goos string = runtime.GOOS
	fmt.Printf("The operating system is: %s\n", goos)
	path := os.Getenv("PATH")
	fmt.Printf("Path is %s\n", path)
}
