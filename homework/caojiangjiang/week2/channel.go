package main

import (
	"fmt"
	"time"
)

//
//无缓冲通道，读数据阻塞
//无缓冲通道，写数据阻塞
//有缓冲通道，读数据阻塞
//有缓冲通道，通道数据已满，写操作阻塞
//

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int, 3)
	ch4 := make(chan int, 3)
	// 无缓冲通道，读数据阻塞
	go noBufferChannelWirteBlock(ch1)
	<-ch1
	fmt.Println("快点吧，我等你等的花都谢了")

	// 无缓冲通道，写数据阻塞
	go noBufferChannelReadBlock(ch2)
	ch2 <- 1
	fmt.Println("快点吧，我等你等的等的花都谢了")

	// 有缓冲通道，读数据阻塞
	go bufferChannelReadBlock(ch3)
	<-ch3
	fmt.Println("快点吧，我等你等的花都谢了花都谢了")

	// 有缓冲通道，通道数据已满，写操作阻塞
	go bufferChannelWriteBlock(ch4)
	time.Sleep(6 * 1e9)
	fmt.Println("我睡了6秒钟")
	<-ch4
}

func noBufferChannelWirteBlock(ch1 chan int) {
	time.Sleep(6 * 1e9)
	fmt.Println("我睡了6秒钟，你等等我哈，Love you, mm ❤️")
	ch1 <- 1
}

func noBufferChannelReadBlock(ch2 chan int) {
	fmt.Println("我you睡了6秒钟，你等等我哈，Love you again, mm ❤️")
	<-ch2
}

func bufferChannelReadBlock(ch3 chan int) {
	time.Sleep(6 * 1e9)
	fmt.Println("我睡了6秒钟，你等等我哈")
	ch3 <- 1
}

func bufferChannelWriteBlock(ch4 chan int) {
	ch4 <- 1
	ch4 <- 2
	ch4 <- 3
	ch4 <- 4
	fmt.Println("终于结束了哈")
}
