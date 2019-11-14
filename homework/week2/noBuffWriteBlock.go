package main

import (
	"fmt"
	"time"
)

func dark(ch chan int)  {
	fmt.Println("its 23:59:59")
	<-ch
}

func brightDay(ch chan int)  {
	ch <- 1
	fmt.Println("its 12:00:00")
}

func morning() {
	fmt.Println("its 07:00:00")
}

func main() {
	/*
	无缓冲通道，写数据阻塞
	打印顺序
	its 07:00:00
	its 23:59:59
	its 12:00:00
	 */
	ch := make(chan int)
	//go morning()
	fmt.Println("its 07:00:00")
	go brightDay(ch)
	go dark(ch)
	time.Sleep(8 * 1e9)
}

