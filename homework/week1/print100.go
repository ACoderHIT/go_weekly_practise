package main

import (
	"fmt"
)

/*
两个协程交替打印数字1-100
 */
func PrintNum(ch chan int, ch2 chan int){
	for i:=1; i<=100; i=i+2 {
		fmt.Println("协程1打印了：", <-ch2)
		ch<-i+1
	}
}
func PrintNum2(ch chan int, ch2 chan int, ch3 chan int){
	for i:=1; i<100; i=i+2 {
		ch2<-i
		fmt.Println("协程2打印了：", <-ch)
	}
	ch3<-1
}

func main() {
	ch:=make(chan int)
	ch2:=make(chan int)
	ch3:=make(chan int)
	go PrintNum(ch, ch2)
	go PrintNum2(ch, ch2, ch3)

	<-ch3
	//time.Sleep(5*1e9)
}

