package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {

	w.Header().Add("WWW-Authenticate", "Basic realm")
	w.Header().Set("WWW-Authenticate", "Basic realm=”google.com”")
	w.Header().Add("WWW-Authenticate", "Basic realm")
	w.WriteHeader(http.StatusUnauthorized)
	fmt.Fprintf(w, "hello\n")
}

func main() {
	http.HandleFunc("/hello", hello)
	http.ListenAndServe(":8080", nil)
}
