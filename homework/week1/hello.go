package main

import (
	"net/http"
	"fmt"
)

/*
使用协程实现一个服务端，要求：
通过浏览器访问响应Hello world
不可以使用snow框架
 */
func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello world")
}

func main() {
	http.HandleFunc("/", helloWorld)
	http.ListenAndServe(":8777", nil)
}
