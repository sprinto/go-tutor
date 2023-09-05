package main

import (
	"fmt"

	"github.com/sprinto/go-tutor/greetings"
)

func main() {
	message := greetings.Hello("Felix")
	fmt.Println(message)
}
