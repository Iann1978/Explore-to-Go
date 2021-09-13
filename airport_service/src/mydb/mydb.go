package mydb

import (
	"database/sql"
	"fmt"
	"os"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func RemoveDb() {
	fmt.Println("removeDb()")
	os.Remove("foo.db")
}

func CreateDb() {
	fmt.Println("CreateDb()")
	db, err := sql.Open("sqlite3", "foo.db")
	checkErr(err)

	stmt, err := db.Prepare("create table userinfo(username text, password text)")
	checkErr(err)

	_, err = stmt.Exec()
	checkErr(err)

	// Insert the data
	stmt, err = db.Prepare("insert into userinfo(username, password) values(?,?)")
	checkErr(err)

	res, err := stmt.Exec("aaa", "bbb")
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)

	fmt.Println("id:", id)
}
