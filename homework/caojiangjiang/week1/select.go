package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	var c chan int

	for {
		select {
		case s := <-c:
			log.Println(s)
		default:
			time.Sleep(100 * time.Millisecond)
			fmt.Println("等待channel数据")
			os.Stdout.Sync()
			return
		}
	}
}
