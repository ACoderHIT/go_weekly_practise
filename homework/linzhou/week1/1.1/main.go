package main

import (
	"fmt"
)

func f(from string) {
	for i := 0; i < 6; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	messages := make(chan int)

	go func() { messages <- 23 }()

	msg := <-messages
	fmt.Println(msg)
}
