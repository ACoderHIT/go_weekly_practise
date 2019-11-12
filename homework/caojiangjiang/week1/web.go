package main

import (
    "fmt"
    "net/http"
    "log"
)

func sayhello(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello world!") // 这个写入到 w 的是输出到客户端的
}

func main() {
    http.HandleFunc("/", sayhello) // 设置访问的路由
    err := http.ListenAndServe(":9090", nil) // 设置监听的端口
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}