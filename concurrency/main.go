package main

import (
	"fmt"
	"time"
)

func pinger(c chan string) {
	for i := 0; i < 10; i++ {
		msg := fmt.Sprintf("%v pinger", i)
		time.Sleep(time.Second * 1)
		c <- msg
	}
	fmt.Println("Pinger done")
}

func printer(c chan string) {
	for {
		msg := <-c
		fmt.Println(msg)
	}
}

func ponger(c chan string) {
	for i := 0; i < 5; i++ {
		msg := fmt.Sprintf("%v ponger", i)
		time.Sleep(time.Second * 2)
		c <- msg
	}
	fmt.Println("Ponger done")
}

func main() {
	c := make(chan string, 1)
	go pinger(c)
	go ponger(c)
	go printer(c)

	var input string
	fmt.Scanln(&input)
}
