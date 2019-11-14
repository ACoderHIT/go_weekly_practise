package main

import (
	"fmt"
)

//可以用e这个channel来控制两个协程严格的顺序
var info = make(chan int, 1)
var e = make(chan int)
func main() {
	go run2(e)
	for {
		<-e
		tp := <-info
		fmt.Println("输出:", tp)
	}
}

func run2(e chan int) {
	for i := 1; ; i++ {
		info<-i
		fmt.Println("输入:", i)
		e<-1
	}
}
