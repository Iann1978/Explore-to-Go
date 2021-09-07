package main

import (
	"encoding/json"
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

type LoginResp struct {
	ErrorString string
	Username    string
}

func login(w http.ResponseWriter, req *http.Request) {

	fmt.Println("Response for login.")

	key, hasKey := getParam(req, "key")

	fmt.Println("hasKey:", hasKey)
	fmt.Println("key:", key)

	username, hasUsername := getParam(req, "username")
	password, hasPassword := getParam(req, "password")
	fmt.Println("hasUsername:", hasUsername)
	fmt.Println("username:", username)
	fmt.Println("hasPassword:", hasPassword)
	fmt.Println("password:", password)

	user, err := users.UserLongin(username, password)
	fmt.Println("user:", user)
	fmt.Println("err:", err)

	loginResp := &LoginResp{ErrorString: "Succeed.", Username: user.Username}
	if err != nil {
		loginResp.ErrorString = err.Error()
	}

	//fmt.Fprintln(w, loginResp)

	jsonLoginResp, _ := json.Marshal(loginResp)

	fmt.Fprintln(w, string(jsonLoginResp))

	if hasUsername && hasPassword {
		fmt.Fprintf(w, "Succeed!\n")
	} else {
		fmt.Fprintf(w, "Failed!\n")
	}

}
