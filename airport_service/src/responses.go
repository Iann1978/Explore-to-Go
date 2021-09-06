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

func getParam(req *http.Request, key string) (string, bool) {
	values, ok := req.URL.Query()[key]

	var nilret string
	if !ok || len(values[0]) < 1 {
		fmt.Println("False in getParam.")
		return nilret, false
	}

	return values[0], true

}

func login(w http.ResponseWriter, req *http.Request) {

	fmt.Println("Response for login.")

	key, hasKey := getParam(req, "key")

	fmt.Println("hasKey:", hasKey)
	fmt.Println("key:", key)

}
