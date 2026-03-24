package main

import (
	"fmt"

	"go_learn/greetings"

	"rsc.io/quote"
)

func main() {
	fmt.Println(quote.Go())

	message := greetings.Hello("5000")
	fmt.Println(message)
}
