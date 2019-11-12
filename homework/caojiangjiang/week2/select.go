package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

//
//自行实现一个常驻进程，读取channel数据，要求如下：
//将读取到的数据，写入日志
//如果channel内没数据，定期刷新等待信息到屏幕
//练习select+定时器
//2019-11-11

func main() {
	var c chan int
	file, err := os.Create("test.log")      // 创建日志文件
	if err != nil {
		fmt.Println("创建日志失败")
	}
	// 创建logger对象　这种方式会显示触发日志文件行数
	logger := log.New(file, "^-------^", log.LstdFlags|log.Llongfile)
	for {
		select {
		case s := <-c:
			logger.Println(s)
		default:
			timer1 := time.NewTimer(10000 * time.Millisecond)
			<-timer1.C
			fmt.Println("等待channel数据")
		}
	}
}
