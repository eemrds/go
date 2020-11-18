package main

func printer(c chan string) {
	for i := 0; ; i++ {
		c <- "test"
	}
}
