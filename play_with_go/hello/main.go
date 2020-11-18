package main

import (
	"fmt"
	"test/greetings"
)

func main() {
	message := greetings.Hello("Erik")
	fmt.Println(message)
}
