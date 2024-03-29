package main

import (
	"airport_service/data"
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

	fmt.Printf("http://%s:8099/hello\n", ipstr)
	fmt.Printf("http://%s:8099/rebuilddb\n", ipstr)
	fmt.Printf("http://%s:8099/regist?username=aaa&&password=bbb\n", ipstr)
	fmt.Printf("http://%s:8099/login?username=aaa&&password=bbb\n", ipstr)
	fmt.Printf("http://%s:8099/logout?username=aaa&&session=bbb\n", ipstr)
	fmt.Printf("http://%s:8099/uploadpos?session=aaa&&longitude=0.0&&latitude=0.0\n", ipstr)
	fmt.Printf("http://%s:8099/pullonlineusers?session=aaa\n", ipstr)
}

func main() {
	aaa := &data.LoginResp2{"aaabbb"}
	fmt.Println(aaa)
	showPromopt()

	users.Open()

	http.HandleFunc("/hello", hello)
	http.HandleFunc("/rebuilddb", rebuilddb)
	//http.HandleFunc("/headers", headers)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)
	http.HandleFunc("/regist", regist)
	http.HandleFunc("/uploadpos", uploadpos)
	http.HandleFunc("/pullonlineusers", pullonlineusers)

	http.ListenAndServe(":8099", nil)

}
