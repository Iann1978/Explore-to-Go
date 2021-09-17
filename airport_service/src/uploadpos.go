package main

import (
	"airport_service/data"
	"airport_service/user"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

type UploadposResp struct {
	ErrorCode   data.ErrorCode
	ErrorString string
	//Session     string
}

func getFloat64Param(req *http.Request, key string) (float64, bool) {
	values, ok := req.URL.Query()[key]

	if !ok || len(values[0]) < 1 {
		fmt.Println("False in getParam.")
		return 0, false
	}

	retval, _ := strconv.ParseFloat(values[0], 64)
	return retval, true
}

func uploadpos(w http.ResponseWriter, req *http.Request) {

	fmt.Println("uploadpos")

	fmt.Println("Response for logout.")
	defer fmt.Println("\n")
	// define variable for responsing
	resp := UploadposResp{ErrorCode: data.NoError, ErrorString: data.NoError.String()}

	// get parameters
	session, hasSession := getParam(req, "session")
	longitude, hasLongitude := getFloat64Param(req, "longitude")
	latitude, hasLatitude := getFloat64Param(req, "latitude")
	fmt.Println("hasSession:", hasSession)
	fmt.Println("session:", session)
	fmt.Println("hasLongitude:", hasLongitude)
	fmt.Println("longitude:", longitude)
	fmt.Println("hasLatitude:", hasLatitude)
	fmt.Println("latitude:", latitude)

	if !hasSession {
		resp.ErrorCode = data.ParameterError
		resp.ErrorString = data.ParameterError.String()
		jsonResp, _ := json.Marshal(resp)
		fmt.Fprintf(w, string(jsonResp))
		return
	}

	// execute uploadpos
	err := users.UploadPos(session, longitude, latitude)
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
