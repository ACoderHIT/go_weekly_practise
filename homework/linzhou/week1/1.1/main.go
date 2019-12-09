package main

import (
	"fmt"
)

func print(count chan int, trigger chan int, done chan bool, prefix string) {
	for {
		odd := <-count
		if odd <= 100 {
			fmt.Println(prefix, " : ", odd)
			trigger <- odd + 1
		} else {
			done <- true
			return
		}
	}
}
func printEven(count chan int, trigger chan int, done chan bool) {
	for {
		odd := <-count
		if odd <= 100 {
			fmt.Println("printEven: ", odd)
			trigger <- odd + 1
		} else {
			done <- true
			return
		}
	}
}

func main() {
	odd := make(chan int)
	even := make(chan int)
	done := make(chan bool)

	go print(odd, even, done, "打印奇数")
	go print(even, odd, done, "打印偶数")
	odd <- 1
	<-done
	fmt.Println("done")
}
