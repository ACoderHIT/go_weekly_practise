package main

import (
	"log"
	"net/http"
)

//问题 使用协程实现一个服务端，通过浏览器访问响应Hello world

func main() {
	http.HandleFunc("/hello", sayHello)
	log.Fatal(http.ListenAndServe(":8262", nil))
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello world!"))
}