package main

import "fmt"


/*
无缓冲通道，读数据阻塞
 */
func main() {
	ch := make(chan int)
	fmt.Println(<-ch)
}
