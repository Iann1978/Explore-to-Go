package main

import (
	"fmt"
	"net/http"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"myfmt1"
	"errors"
	"net"
	"user"
)

var users user.UserDatabase

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello\n")
}

func headers(w http.ResponseWriter, req *http.Request) {

	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func login(w http.ResponseWriter, req *http.Request) {

	keys, ok := req.URL.Query()["key"]

	if !ok || len(keys[0]) < 1 {
		fmt.Fprintf(w,"Url Param 'key' is missing!!!")
		return
	}

	key := keys[0]
	fmt.Fprintf(w, "key: %v\n", key)

	usernames, ok := req.URL.Query()["username"]
	if !ok || len(usernames[0]) < 1 {
		fmt.Fprintf(w, "Url Param 'username' is missing!!!")
		return
	}
	req_username := usernames[0]
	fmt.Fprintf(w, "username: %v\n", req_username)

	userid, err := users.HasUser(req_username)
	checkErr(err)

	fmt.Println("userid:", userid)

	passwords, ok := req.URL.Query()["password"]
	if !ok || len(passwords[0]) < 1 {
		fmt.Fprintf(w, "Url Param 'password' is missing!!!")
		return
	}
	req_password := passwords[0]
	fmt.Fprintf(w, "password: %v\n", req_password)


	db, err := sql.Open("sqlite3", "./foo.db")
	checkErr(err)


	rows, err := db.Query("select username, password from userinfo")
	checkErr(err)

	
	for rows.Next() {
		var username string
		var password string

		err = rows.Scan(&username, &password)
		checkErr(err)
		fmt.Fprintln(w, "username in db:" , username)
		fmt.Fprintln(w, "password in db:", password)
	}

	fmt.Fprintf(w, "Succeed!")
}

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

func main() {

	addrs, err := net.InterfaceAddrs()
	checkErr(err)

	for _, address := range addrs {

		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				fmt.Println(ipnet.IP.String())
			}
		}
	}

	ipstr, err := getClientIp()
	checkErr(err)

	err = users.Open()
	checkErr(err)

	myfmt1.Println()
	fmt.Println("http://", ipstr, ":8099/hello")
	fmt.Println("http://", ipstr, ":8099/login")
	fmt.Println("http://", ipstr, ":8099/login?key=123&username=aaa&&password=bbb")
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)
	http.HandleFunc("/login", login)

	http.ListenAndServe(":8099", nil)
}
