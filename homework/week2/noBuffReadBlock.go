package main

import "fmt"

func getOut(ch chan int)  {
	for i:=1; i < 10; i ++ {
		ch <- i
	}
}

func main() {
	/*
	无缓冲通道，读数据阻塞
	结果：只打印了 1
	而不是打印 1-10
	 */
	ch := make(chan int)
	go getOut(ch)
	fmt.Println(<-ch)

}
