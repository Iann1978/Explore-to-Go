package main

import (
	"database/sql"
	"flag"
	"fmt"
	"os"
	"path/filepath"

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

	filename := filepath.Join(*path, "foo.db")

	fmt.Println("filename: ", filename)

	if *removeDb {
		os.Remove(filename)
	}

	db, err := sql.Open("sqlite3", filename)
	checkErr(err)

	stmt, err := db.Prepare("create table userinfo(username text, password text, session text)")
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

	fmt.Println(id)
}
