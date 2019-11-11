package main

import "fmt"

//问题
//- 自行实现一个常驻进程，读取channel数据，要求如下：
//- 将读取到的数据，写入日志
//- 如果channel内没数据，定期刷新等待信息到屏幕

//阻塞的方式
//func main() {
//	e := make(chan int)
//	go run1(e)
//	for {
//		<-e
//	}
//}
//
//
//func run1(e chan int) {
//	for {
//		e<-1
//	}
//}

//非阻塞的方式
//当第一个消息入channel之后，在第一个出之前，第二个不会进去 应该是同样可以实现的
//不管缓冲区有多大，都可以用e的这个channel来控制两个协程来交替进行输入输出
func main() {
	e := make(chan int, 1)
	go run2(e)
	for {
		<-e
		fmt.Println("输出")
	}
}


func run2(e chan int) {
	for {
		e<-1
		fmt.Println("输入")
	}
}
