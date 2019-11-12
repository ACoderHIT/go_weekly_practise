package main

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"os/signal"
	"time"
)

var logChan  = make(chan map[string]interface{})

var requestStatusMap = map[int]bool{}


var done = make(chan bool, 1)
var quit = make(chan os.Signal, 1)

//在第一周接收请求作业基础上，要求如下：
//将接收到请求相关的业务日志异步方式写入文件
//接收系统关闭指令，并可以做到平滑关闭



//为什么这样可以平滑重启？（待确认正确性）
// 正常情况下是server.ListenAndServe() 这个位置hang住整个进程的
// 可以把这个程序看成两部分，1个是web服务的监听部分，一个是处理部分， 如果web服务器不开启了，那么就不能处理新进来的请求了（可以理解为一个带路的）
// 真正让这个请求断掉  是因为主进程（main）被kill
// 所以平滑重启的原理就是，先kill掉web服务器，不让新的请求进来，等现有的全部请求完了，然后结束当前进程
func main() {
	server := newServer()
	signal.Notify(quit, os.Interrupt)
	go monitorKill(server, quit)
	server.ListenAndServe()
	<-done
}


func newServer() *http.Server {
	router := http.NewServeMux()
	router.HandleFunc("/hello", sayHello)
	return &http.Server{
		Addr:         ":8262",
		Handler:      router,
	}
}

func monitorKill(server *http.Server, quit <-chan os.Signal)  {
	<-quit
	go shutDown(server)
	for {
		if len(requestStatusMap) != 0 {
			fmt.Println("目前还有进行中的请求，请稍等")
			time.Sleep(time.Second * 1)
			continue
		} else {
			close(done)
			break
		}
	}
}

func shutDown(server *http.Server) {
	if err := server.Shutdown(context.Background()); err != nil {
		fmt.Println(err)
	}
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	go WriteInfo()//请求写日志
	var uniqueId = GenerateRangeNum(1, 1000)
	requestStatusMap[uniqueId] = false
	url := r.URL.Path
	query  := r.URL.RawQuery
	method := r.Method
	a := map[string] interface{}{
		"url" : url,
		"method" : method,
		"query" : query,
		"response": "hello world!",
	}
	logChan<-a
	w.Write([]byte("hello world!"))
	time.Sleep(time.Second * 10)
	delete(requestStatusMap, uniqueId)
}

func WriteInfo() {
	info := <-logChan
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
		infostr, _ := json.Marshal(info)
		_,err=f.WriteAt([]byte(string(infostr) +"\n"), n)
	}
}

func GenerateRangeNum(min int, max int) int {
	if min == max {
		return min
	}
	rand.Seed(time.Now().Unix())
	randNum := rand.Intn(max-min) + min
	return randNum
}
