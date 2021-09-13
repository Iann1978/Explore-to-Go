package main

import (
	"fmt"
	"net/http"
)

func regist(w http.ResponseWriter, req *http.Request) {

	username, hasUsername := getParam(req, "username")
	password, hasPassword := getParam(req, "password")
	fmt.Println("hasUsername:", hasUsername)
	fmt.Println("username:", username)
	fmt.Println("hasPassword:", hasPassword)
	fmt.Println("password:", password)

	user, err := users.Regist(username, password)
	fmt.Println("user:", user)
	fmt.Println("err:", err)

	if err != nil {
		fmt.Fprintf(w, err.Error())
		fmt.Fprintf(w, "\nFailed!\n")
	} else {
		fmt.Fprint(w, "Succeed!\n")
	}

}
