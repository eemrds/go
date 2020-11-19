package main

import (
	"fmt"
	"log"
	"go/play_with_go/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(1)
	people := []string{"erik", "", "tormod"}
	message, err := greetings.Hellos(people)
	for i, greet := range message{
		if err[i] != nil{
			log.Printf("%v : %v", greet, err[i])
		}
	fmt.Println(message[i])
	}
}
