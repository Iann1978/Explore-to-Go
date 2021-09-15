package main

import (
	"airport_service/data"
	"airport_service/user"
	"encoding/json"
	"fmt"
	"net/http"
)

type RegistResp struct {
	ErrorCode   data.ErrorCode
	ErrorString string
}
type LoginResp1 struct {
	ErrorString string
	//Username    string
	//user        *user.User
}

func regist(w http.ResponseWriter, req *http.Request) {

	// define variable for responsing
	resp := RegistResp{ErrorCode: data.NoError, ErrorString: data.NoError.String()}

	// get parameters
	username, hasUsername := getParam(req, "username")
	password, hasPassword := getParam(req, "password")
	fmt.Println("hasUsername:", hasUsername)
	fmt.Println("username:", username)
	fmt.Println("hasPassword:", hasPassword)
	fmt.Println("password:", password)

	if !hasPassword || !hasPassword {
		resp.ErrorCode = data.ParameterError
		resp.ErrorString = data.ParameterError.String()
		jsonResp, _ := json.Marshal(resp)
		fmt.Fprintf(w, string(jsonResp))
		return
	}

	// execute regist operation
	auser, err := users.Regist(username, password)
	fmt.Println("user:", auser)
	fmt.Println("err:", err)
	if err != nil {
		aa := &user.UserSetError{data.NoError, ""}
		fmt.Println(aa)
		if userSetError, ok := err.(*user.UserSetError); ok {
			resp.ErrorCode = userSetError.ErrorCode
			resp.ErrorString = userSetError.ErrorString
		} else {
			resp.ErrorCode = data.UnknownError
			resp.ErrorString = data.UnknownError.String()
		}
		jsonResp, _ := json.Marshal(resp)
		fmt.Fprintf(w, string(jsonResp))
		//fmt.Fprintf(w, err.Error())
		//fmt.Fprintf(w, "\nFailed!\n")
		return
	}

	// return succeed
	resp.ErrorCode = data.NoError
	resp.ErrorString = data.NoError.String()
	jsonResp, _ := json.Marshal(resp)
	fmt.Fprintf(w, string(jsonResp))
}
