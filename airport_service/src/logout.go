package main

import (
	"airport_service/data"
	"airport_service/user"
	"encoding/json"
	"fmt"
	"net/http"
)

type LogoutResp struct {
	ErrorCode   data.ErrorCode
	ErrorString string
	//Session     string
}

func logout(w http.ResponseWriter, req *http.Request) {

	fmt.Println("Response for logout.")
	defer fmt.Println("\n")
	// define variable for responsing
	resp := LogoutResp{ErrorCode: data.NoError, ErrorString: data.NoError.String()}

	// get parameters
	username, hasUsername := getParam(req, "username")
	//password, hasPassword := getParam(req, "password")
	fmt.Println("hasUsername:", hasUsername)
	fmt.Println("username:", username)
	// fmt.Println("hasPassword:", hasPassword)
	// fmt.Println("password:", password)

	if !hasUsername {
		resp.ErrorCode = data.ParameterError
		resp.ErrorString = data.ParameterError.String()
		jsonResp, _ := json.Marshal(resp)
		fmt.Fprintf(w, string(jsonResp))
		return
	}

	// execute login
	err := users.UserLogout(username)
	//	fmt.Println("user:", user)
	fmt.Println("err:", err)

	if err != nil {
		if userSetError, ok := err.(*user.UserSetError); ok {
			resp.ErrorCode = userSetError.ErrorCode
			resp.ErrorString = userSetError.ErrorString
		} else {
			resp.ErrorCode = data.UnknownError
			resp.ErrorString = data.UnknownError.String()
		}
		jsonResp, _ := json.Marshal(resp)
		fmt.Fprintf(w, string(jsonResp))
		return
	}

	// return succeed
	resp.ErrorCode = data.NoError
	resp.ErrorString = data.NoError.String()
	jsonResp, _ := json.Marshal(resp)
	fmt.Fprintf(w, string(jsonResp))
}
