package main

import (
	"fmt"
	"sync"
)

var (
	w sync.WaitGroup
	v = 0
	)

//问题 两个协程交替打印数字1-100

//思路：两个协程用共享内存v 和 w来进行通信
// w用来hold main防止协程运行完之前main退出
// v则用来协调两个协程轮换打印奇偶数

//方法1 使用共享资源来阻塞
//func main()  {
//	w.Add(2)
//	go print(1);
//	go print(0);
//	w.Wait()
//}
//
//func print(a int) {
//	for i:= 1; i <= 50; {
//		if v == a {continue}
//		fmt.Println("协程id:" + strconv.Itoa(a) + "    " + strconv.Itoa(i*2 - a))
//		v = v ^ 1
//		i++
//	}
//	w.Done()
//}


//方法2 使用无缓冲channel来阻塞


//func main()  {
//	foochannel := make(chan int)
//	total := make(chan int)
//	go print1(foochannel, total);
//	go print2(foochannel);
//	foochannel<-1
//	total<-1
//}
//func print2(c chan int) {
//	for i:= 1; i <= 50; {
//		<-c
//		fmt.Println(i*2-1)
//		i++
//		c<-1
//	}
//}
//func print1(c chan int, b chan int) {
//	time.Sleep(time.Second * 1)
//	for i:= 1; i <= 50; {
//		<-c
//		fmt.Println(i*2)
//		if i != 50 {
//			c<-1
//		}
//		i++
//	}
//	<-b
//}
//



// 不要求goroutine的执行顺序仍然可以行的版本
func main()  {
	c := make(chan int)
	total := make(chan int)
	go print1(c, total);
	go print2(c, total);
	c<-1
	total<-1
}
func print1(c chan int, total chan int) {
	for {
		a := <-c
		fmt.Println(a)
		if a >= 100 {
			<-total
			return
		}
		c<-a+1
	}
}
func print2(c chan int, total chan int) {
	for {
		a := <-c
		fmt.Println(a)
		if a >= 100 {
			<-total
			return
		}
		c<-a+1
	}
}