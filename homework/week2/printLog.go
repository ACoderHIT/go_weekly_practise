package main

import (
	"fmt"
	"time"
	"math/rand"
)

/*
自行实现一个常驻进程，读取channel数据，要求如下：
将读取到的数据，写入日志
如果channel内没数据，定期刷新等待信息到屏幕
 */
func randomNum(ch chan int){
	for {
		time.Sleep(3*1e9)
		ch<-rand.Int()
	}
}

func main() {
	ch := make(chan int)
	go  randomNum(ch)
	timer2 := time.NewTicker(2 * time.Second)
	for {
		select {
		case v :=  <-ch:
			fmt.Println("这是channel的值:", v)
		default:
			<-timer2.C
			fmt.Println("这是要记录日志的地方")
		}
	}
}
