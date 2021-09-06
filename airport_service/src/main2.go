package main

import (
	"airport_service/myfmt1"
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "hello\n")
}
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
func main() {
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
