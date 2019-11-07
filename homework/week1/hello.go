package main

import (
	"net/http"
	"fmt"
)

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello world")
}

func main() {
	http.HandleFunc("/", helloWorld)
	http.ListenAndServe(":8777", nil)
}
