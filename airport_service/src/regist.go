package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ErrorCode int32

const (
	NoError        ErrorCode = 0
	UnknownError   ErrorCode = 1
	ParameterError ErrorCode = 2
	UserExist                = 3
	UserNotExist             = 4
)

func (e ErrorCode) String() string {
	switch e {
	case NoError:
		return "No Error"
	case UnknownError:
		return "Unknown Error"
	case ParameterError:
		return "Parameter Error"
	case UserExist:
		return "User Already Exist"
	case UserNotExist:
		return "User Not Exist"
	}
	return "unknown"
}

type RegistResp struct {
	ErrorCode   ErrorCode
	ErrorString string
}
type LoginResp1 struct {
	ErrorString string
	//Username    string
	//user        *user.User
}

func regist(w http.ResponseWriter, req *http.Request) {

	// define variable for responsing
	resp := RegistResp{ErrorCode: NoError, ErrorString: NoError.String()}

	// get parameters
	username, hasUsername := getParam(req, "username")
	password, hasPassword := getParam(req, "password")
	fmt.Println("hasUsername:", hasUsername)
	fmt.Println("username:", username)
	fmt.Println("hasPassword:", hasPassword)
	fmt.Println("password:", password)

	if !hasPassword || !hasPassword {
		resp.ErrorCode = ParameterError
		resp.ErrorString = ParameterError.String()
		jsonResp, _ := json.Marshal(resp)
		fmt.Fprintf(w, string(jsonResp))
		return
	}

	// execute regist operation
	user, err := users.Regist(username, password)
	fmt.Println("user:", user)
	fmt.Println("err:", err)
	if err != nil {
		resp.ErrorCode = UnknownError
		resp.ErrorString = UnknownError.String()
		jsonResp, _ := json.Marshal(resp)
		fmt.Fprintf(w, string(jsonResp))
		//fmt.Fprintf(w, err.Error())
		//fmt.Fprintf(w, "\nFailed!\n")
		return
	}

	// return succeed
	resp.ErrorCode = NoError
	resp.ErrorString = NoError.String()
	jsonResp, _ := json.Marshal(resp)
	fmt.Fprintf(w, string(jsonResp))
}
