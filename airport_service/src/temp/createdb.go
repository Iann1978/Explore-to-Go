package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

func main() {

	db, err := sql.Open("sqlite3", "./foo.db")
	checkErr(err)

	// Create table userinfo
	stmt, err := db.Prepare("create table userinfo(username text, password text)")
	checkErr(err)

	_, err = stmt.Exec()
	checkErr(err)

	stmt, err = db.Prepare("insert into userinfo(username, password) values(\"test\", \"123\")")
	checkErr(err)

	_, err = stmt.Exec()
	checkErr(err)

	db.Close()

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
