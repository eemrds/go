package main

import (
	"fmt"

	greetings "github.com/eemrds/go/pay_with_go/greetings"
)

func main() {
	message := greetings.Hello("Erik")
	fmt.Println(message)
}
