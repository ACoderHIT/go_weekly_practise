package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", func(resp http.ResponseWriter, req *http.Request) {
		resp.Write([]byte("hello world"))

	})
	http.ListenAndServe(":9527", nil)
}
