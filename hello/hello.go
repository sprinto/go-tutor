package main

import (
	"fmt"

	"example.com/greetings"
)

func main() {
	message := greetings.Hello("Felix")
	fmt.Println(message)
}
