package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

//问题
//- 自行实现一个常驻进程，读取channel数据，要求如下：
//- 将读取到的数据，写入日志
//- 如果channel内没数据，定期刷新等待信息到屏幕



func main() {
	c := make(chan int)
	run(c)
}

func timer(boo <-chan time.Time, c chan int) {
	for {
		select {
			case <-boo:
				c<-rand.Int()
				boo = time.After(5000 * time.Millisecond)
			default:
				time.Sleep(1000 * time.Millisecond)
				fmt.Println("赋值倒计时")
		}
	}
}

func run (c chan int) {
	boo := time.After(5000 * time.Millisecond)
	go timer(boo, c)
	for {
		select {
			case a := <-c:
				writeInfo(strconv.Itoa(a))
				fmt.Println(strconv.Itoa(a) + "信息已写入日志")
			default:
				time.Sleep(100 * time.Millisecond)
		}
	}
}

func writeInfo(s string) {
	fileName := "/tmp/weekhomework.log"
	_, err := os.Stat(fileName)
	if err != nil || os.IsNotExist(err) {
		_, _ = os.Create(fileName)
	}
	f,err := os.OpenFile(fileName, os.O_WRONLY, 0644)
	defer f.Close()
	if err !=nil {
		fmt.Println(err.Error())
	} else {
		//追加写入   为什么O_APPEND 模式无法写入？ todo
		n, _ := f.Seek(0, 2)
		_,err=f.WriteAt([]byte(s+"\n"), n)
	}
}
