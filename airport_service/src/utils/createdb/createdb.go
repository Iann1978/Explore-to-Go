package main

import (
	"database/sql"
	"flag"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	fmt.Println("create db")

	removeDb := flag.Bool("removeDb", false, "Remove Database file if exists")
	path := flag.String("path", "./", "The path of db file to operation.")

	flag.Parse()

	fmt.Println("removeDb: ", *removeDb)
	fmt.Println("path: ", *path)

	if *removeDb {
		os.Remove("foo.db")
	}

	db, err := sql.Open("sqlite3", "./foo.db")
	checkErr(err)

	stmt, err := db.Prepare("create table userinfo(username text, password text)")
	checkErr(err)

	_, err = stmt.Exec()
	checkErr(err)

	// Insert the data
	stmt, err = db.Prepare("insert into userinfo(username, password) values(?,?)")
	checkErr(err)

	res, err := stmt.Exec("astaxie", "Developer Department")
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)

	fmt.Println(id)
}
