package main

import (
	"airport_service/data"
	"airport_service/user"
	"encoding/json"
	"fmt"
	"net/http"
)

type PullOnlineUsersResp struct {
	ErrorCode   data.ErrorCode
	ErrorString string
	Users       []user.User
	//Session     string
}

func pullonlineusers(w http.ResponseWriter, req *http.Request) {

	//fmt.Fprintln(w, "pullonlineusers")

	fmt.Println("Response for pullonlineusers.")
	defer fmt.Println("\n")
	// define variable for responsing
	resp := PullOnlineUsersResp{ErrorCode: data.NoError, ErrorString: data.NoError.String()}

	// get parameters
	session, hasSession := getParam(req, "session")
	// longitude, hasLongitude := getFloat64Param(req, "longitude")
	// latitude, hasLatitude := getFloat64Param(req, "latitude")
	fmt.Println("hasSession:", hasSession)
	fmt.Println("session:", session)
	// fmt.Println("hasLongitude:", hasLongitude)
	// fmt.Println("longitude:", longitude)
	// fmt.Println("hasLatitude:", hasLatitude)
	// fmt.Println("latitude:", latitude)

	if !hasSession {
		resp.ErrorCode = data.ParameterError
		resp.ErrorString = data.ParameterError.String()
		jsonResp, _ := json.Marshal(resp)
		fmt.Fprintf(w, string(jsonResp))
		return
	}

	// execute uploadpos
	onlineUsers, err := users.AllUsers(session)
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

	resp.Users = onlineUsers
	for _, usr := range onlineUsers {
		fmt.Println(usr)
	}

	// return succeed
	resp.ErrorCode = data.NoError
	resp.ErrorString = data.NoError.String()
	jsonResp, _ := json.Marshal(resp)
	fmt.Fprintf(w, string(jsonResp))

}
