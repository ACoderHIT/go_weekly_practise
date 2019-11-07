package main

import (
	"fmt"
	"time"
	"math/rand"
)

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
