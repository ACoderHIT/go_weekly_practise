package main

import (
	"sync"
)

var (
	w sync.WaitGroup
	v = 0
	)

//问题
//- 无缓冲通道，读数据阻塞
//- 无缓冲通道，写数据阻塞
//- 有缓冲通道，读数据阻塞
//- 有缓冲通道，通道数据已满，写操作阻塞


func main()  {
	readWithoutBuffer()
	//all goroutines are asleep - deadlock!
	writeWithoutBuffer()
	//all goroutines are asleep - deadlock!
	readWithBuffer()
	//all goroutines are asleep - deadlock!
	writeWithBufferNotDeadlock()
	//有缓冲区 且没满 正常写入
	writeWithBuffer()
	//all goroutines are asleep - deadlock!
}

func readWithoutBuffer() {
	c := make(chan int)
	<-c
	close(c)
}

func writeWithoutBuffer() {
	c := make(chan int)
	c<-1
	close(c)
}

func readWithBuffer() {
	c := make(chan int, 1)
	<-c
	close(c)
}

func writeWithBufferNotDeadlock() {
	c := make(chan int, 1)
	c<-1
	close(c)
}

func writeWithBuffer() {
	c := make(chan int, 1)
	c<-1
	c<-1
	close(c)
}

