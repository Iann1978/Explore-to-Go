package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type LoginResp struct {
	ErrorCode   ErrorCode
	ErrorString string
	Session     string
}

func login(w http.ResponseWriter, req *http.Request) {

	fmt.Println("Response for login.")
	defer fmt.Println("\n")
	// define variable for responsing
	resp := LoginResp{ErrorCode: NoError, ErrorString: NoError.String()}

	// get parameters
	username, hasUsername := getParam(req, "username")
	password, hasPassword := getParam(req, "password")
	fmt.Println("hasUsername:", hasUsername)
	fmt.Println("username:", username)
	fmt.Println("hasPassword:", hasPassword)
	fmt.Println("password:", password)

	if !hasUsername || !hasPassword {
		resp.ErrorCode = ParameterError
		resp.ErrorString = ParameterError.String()
		jsonResp, _ := json.Marshal(resp)
		fmt.Fprintf(w, string(jsonResp))
		return
	}

	// execute login
	user, err := users.UserLongin(username, password)
	fmt.Println("user:", user)
	fmt.Println("err:", err)

	if err != nil {
		resp.ErrorCode = UnknownError
		resp.ErrorString = UnknownError.String()
		jsonResp, _ := json.Marshal(resp)
		fmt.Fprintf(w, string(jsonResp))
		return
	}

	// return succeed
	resp.ErrorCode = NoError
	resp.ErrorString = NoError.String()
	jsonResp, _ := json.Marshal(resp)
	fmt.Fprintf(w, string(jsonResp))

}
