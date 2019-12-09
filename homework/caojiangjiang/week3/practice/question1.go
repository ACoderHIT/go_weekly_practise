package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"log"
	"os"
	"strings"
	"time"
)

func sayhello(w http.ResponseWriter, r *http.Request) {
	go writeLog(w, r);
	fmt.Fprintf(w, "Hello world!")     // 这个写入到 w 的是输出到客户端的
}

/**
 * 方法功能：监听端口，异步写入日志
 */
func main() {
	http.HandleFunc("/", sayhello) // 设置访问的路由
	err := http.ListenAndServe(":9090", nil) // 设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func writeLog(w http.ResponseWriter, r *http.Request)  {
	url := r.URL.Path
	query  := r.URL.RawQuery
	method := r.Method
	methodInfoMap := map[string] interface{}{
		"url" : url,
		"method" : method,
		"query" : query,
		"response": "hello world!",
	}
	methodInfoJson, _ := json.Marshal(methodInfoMap)
	methodInfoStr := string(methodInfoJson)
	fileName := "/tmp/homework.log"        // 根目录下的tmp目录下
	_, err := os.Stat(fileName)
	if err != nil || os.IsNotExist(err)  {
		os.Create(fileName)
	}
	fd,_:=os.OpenFile(fileName,os.O_RDWR|os.O_CREATE|os.O_APPEND,0666)
	defer fd.Close()
	fd_time:=time.Now().Format("2006-01-02 15:04:05");
	fd_content:=strings.Join([]string{"======",fd_time,"=====",methodInfoStr,"\n"},"")
	buf:=[]byte(fd_content)
	fd.Write(buf)
}