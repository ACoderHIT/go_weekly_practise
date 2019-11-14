package main

import "fmt"

/*
fatal error: all goroutines are asleep - deadlock!
 */
func main() {
	happy := make(chan bool, 5)
	for i := 0; i < 10; i++ {
		happy<-true
		fmt.Println("你笑起来真好看")
	}
	for i := 0; i < 10; i++ {
		<-happy
	}
}