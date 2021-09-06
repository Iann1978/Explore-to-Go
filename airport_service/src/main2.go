package main

import (
	"airport_service/myfmt1"
	"airport_service/user"
	"database/sql"
	"errors"
	"fmt"
	"net"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

var users user.UserDatabase

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello\n")
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

func showPromopt() {

	ipstr, err := getClientIp()
	checkErr(err)

	fmt.Println("http://", ipstr, ":8099/hello")
	fmt.Println("http://", ipstr, ":8099/login")
	fmt.Println("http://", ipstr, ":8099/login?key=123&username=aaa&&password=bbb")

}

func main() {

	showPromopt()

	fmt.Println(users)
	users.Open()
	fmt.Println(users)
	fmt.Println("haha")
	myfmt1.Println()
	myprintf()

	db, err := sql.Open("sqlite3", "./foo.db")
	checkErr(err)

	// stmt, err := db.Prepare("create table userinfo(username text, password text)")
	// checkErr(err)

	// _, err = stmt.Exec()
	// checkErr(err)

	// Insert the data
	stmt, err := db.Prepare("insert into userinfo(username, password) values(?,?)")
	checkErr(err)

	res, err := stmt.Exec("astaxie", "Developer Department")
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)

	fmt.Println(id)

	http.HandleFunc("/hello", hello)
	//http.HandleFunc("/headers", headers)
	//http.HandleFunc("/login", login)

	http.ListenAndServe(":8099", nil)
}
