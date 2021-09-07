package user

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	userid   int32
	Username string
	x        float64
	y        float64
}

type UserDatabase struct {
	db *sql.DB
}

type UserSet interface {
	UserLongin(username string, password string) (*User, error)
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

func (users *UserDatabase) UserLongin(username string, password string) (*User, error) {
	user := &User{userid: 0, Username: username, x: 0, y: 0}

	return user, nil

}
