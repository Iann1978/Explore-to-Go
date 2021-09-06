package user

import (
	"database/sql"
	"fmt"
)

type User struct {
	userid int32
	name   string
	x      float64
	y      float64
}

type UserDatabase struct {
	db *sql.DB
}

type UserSet interface {
	HasUser(username string) (int32, error)
}

func (users *UserDatabase) Open() error {
	fmt.Println("UserDatabase.Open()")
	db, err := sql.Open("sqlite3", "./foo.db")
	if db != nil {
		fmt.Println("db != nil")
		users.db = db
	}
	if err != nil {
		fmt.Println(err)
	}
	return err
}

func (users UserDatabase) HasUser(username string) (int32, error) {
	return 0, nil
}
