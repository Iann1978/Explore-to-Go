package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello\n")
}

func headers(w http.ResponseWriter, req *http.Request) {

	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func login(w http.ResponseWriter, req *http.Request) {

	keys, ok := req.URL.Query()["key"]

	if !ok || len(keys[0]) < 1 {
		fmt.Fprintf(w,"Url Param 'key' is missing!!!")
		return
	}

	key := keys[0]
	fmt.Fprintf(w, "key: %v\n", key)

	usernames, ok := req.URL.Query()["username"]
	if !ok || len(usernames[0]) < 1 {
		fmt.Fprintf(w, "Url Param 'username' is missing!!!")
		return
	}
	username := usernames[0]
	fmt.Fprintf(w, "username: %v\n", username)


	passwords, ok := req.URL.Query()["password"]
	if !ok || len(passwords[0]) < 1 {
		fmt.Fprintf(w, "Url Param 'password' is missing!!!")
		return
	}
	password := passwords[0]
	fmt.Fprintf(w, "password: %v\n", password)

	fmt.Fprintf(w, "Succeed!")
}


func main() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)
	http.HandleFunc("/login", login)

	http.ListenAndServe(":8099", nil)
}
