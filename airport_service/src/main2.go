package main

import (
	"airport_service/user"
	"errors"
	"fmt"
	"net"
	"net/http"
)

var users user.UserDatabase

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func getClientIp() (string, error) {
	addrs, err := net.InterfaceAddrs()

	if err != nil {
		return "", err
	}

	for _, address := range addrs {

		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String(), err
			}
		}
	}

	return "", errors.New("Can not find the client ip address!")
}

func showPromopt() {

	ipstr, err := getClientIp()
	checkErr(err)

	fmt.Println("http://", ipstr, ":8099/hello")
	fmt.Println("http://", ipstr, ":8099/login")
	fmt.Println("http://", ipstr, ":8099/login?key=123&username=aaa&&password=bbb")

}

func main() {

	showPromopt()

	users.Open()

	http.HandleFunc("/hello", hello)
	//http.HandleFunc("/headers", headers)
	http.HandleFunc("/login", login)
	http.HandleFunc("/regist", regist)

	http.ListenAndServe(":8099", nil)
}
