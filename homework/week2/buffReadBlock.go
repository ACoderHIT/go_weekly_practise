package main

import "fmt"

/*
打印出的次数在1-10之间
 */
func main() {
	happy := make(chan bool)
	for i := 0; i < 10; i++ {
		go func(i int) {
			happy <- true
			fmt.Println("你笑起来真好看")
		}(i)
	}
	for i := 0; i < 10; i++ {
		<-happy
	}
}