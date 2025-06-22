package main

import (
	"fmt"

	library "github.com/agent3892/library"
)

func main() {
	fmt.Println("Hello, Mr. Go!")

	message := library.Greet()
	fmt.Println(message)
}
