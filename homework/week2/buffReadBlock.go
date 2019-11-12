package main

/*
有缓冲读阻塞
 */
func main() {
	ch := make(chan int, 5)
	ch <- 1
	ch <- 2
	ch <- 3
	ch <- 6
	ch <- 7
	ch <- 7

}