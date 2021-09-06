package main

import (
	"fmt"
	"net/http"
)

func myprintf() {
	fmt.Println("hehe")
}

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello\n")
}
