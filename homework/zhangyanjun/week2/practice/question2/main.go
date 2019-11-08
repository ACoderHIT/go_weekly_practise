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
	t := time.NewTicker(time.Second*1)
	go run(c)
	for {
		select {
			case a :=  <-c:
				fmt.Println("拿到数据")
				writeInfo(strconv.Itoa(a))
			default:
				<-t.C
				fmt.Println("获取数据倒计时")
		}
	}
}

//模拟时不时的有数据进来
func run( c chan int) {
	for {
		time.Sleep(3000 * time.Millisecond)
		c<-rand.Int()
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
