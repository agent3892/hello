package mypackage

import (
	"fmt"

	"github.com/agent3892/hello/mypackage"
)

func main() {
	fmt.Println("Hello, Go!")

	message := mypackage.Greet()
	fmt.Println(message)

}
