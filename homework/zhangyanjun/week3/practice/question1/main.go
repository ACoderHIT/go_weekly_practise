package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

var logChan  = make(chan map[string]interface{})

var requestStatusMap = map[int]bool{}


var done = make(chan bool, 1)
var quit = make(chan os.Signal, 1)

//在第一周接收请求作业基础上，要求如下：
//将接收到请求相关的业务日志异步方式写入文件

func main() {
	server := newServer()
	server.ListenAndServe()
}


func newServer() *http.Server {
	router := http.NewServeMux()
	router.HandleFunc("/hello", sayHello)
	return &http.Server{
		Addr:         ":8262",
		Handler:      router,
	}
}


func sayHello(w http.ResponseWriter, r *http.Request) {
	go WriteInfo()//请求写日志
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
}

func WriteInfo()  {
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
