package main

import (
	"fmt"

	"github.com/sprinto/go-tutor/stringmod/quotes"
	"github.com/sprinto/go-tutor/stringmod/strings"
)

func main() {
	o, e := strings.CountOddEven("12345")
	fmt.Println(o, e) // 3 2
	fmt.Println(quotes.GetEmoji("turtle"))
}
