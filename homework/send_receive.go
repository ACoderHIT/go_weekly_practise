package main

import "fmt"

/*
加练题目：
利用协程向主进程发送数据，要求如下：
协程发送一条数据，主进程接收一条数据，交替进行，不可连续发送，或者连续接收
发送数据的时候输出 send data
接收数据的时候输出 receive data
 */

func test(ch1 chan string, ch2 chan string) {
	for {
		ch1 <- "send data"
		re := <-ch2
		fmt.Println(re)
	}
}

func main()  {
	ch1 := make(chan string)
	ch2 := make(chan string)
	go test(ch1, ch2)
	for {
		re := <-ch1
		fmt.Println(re)
		ch2<-"receive data"
	}
}
