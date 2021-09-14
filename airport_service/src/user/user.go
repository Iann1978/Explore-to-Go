package user

import (
	"database/sql"
	"errors"
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
	db          *sql.DB
	onlineUsers map[int]*User
	counter     int32
}

type UserSet interface {
	Regist(username string, password string) (*User, error)
	UserLongin(username string, password string) (*User, error)
	HasUser(username string) (int32, error)
}

func (users *UserDatabase) Open() error {
	fmt.Println("UserDatabase.Open()")
	db, err := sql.Open("sqlite3", "./foo.db")
	if db != nil {
		fmt.Println("db != nil")
		users.db = db
		users.onlineUsers = make(map[int]*User)
		users.counter = 0
	}
	if err != nil {
		fmt.Println(err)
	}
	return err
}

func (users *UserDatabase) Close() error {
	fmt.Println("UserDatabase.Close()")
	users.db.Close()
	users.onlineUsers = nil
	users.counter = 0
	return nil
}

func (users UserDatabase) HasUser(username string) (int32, error) {
	return 0, nil
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func (users *UserDatabase) UserLongin(username string, userPassword string) (*User, error) {

	rows, err := users.db.Query("select username, password from userinfo where username=?", username)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	user := &User{userid: 0, Username: username, x: 0, y: 0}

	for rows.Next() {

		var username string
		var password string
		if err := rows.Scan(&username, &password); err != nil {
			return nil, err
		} else {
			user.Username = username

			if userPassword == password {
				users.counter++
				user.userid = users.counter
				return user, nil
			} else {
				return nil, errors.New("Password error!!!")
			}

		}
	}

	return user, nil
}

func (users *UserDatabase) Regist(username string, password string) (*User, error) {

	rows, err := users.db.Query("select username, password from userinfo where username=?", username)
	defer rows.Close()
	if err != nil {
		return nil, err
	}

	//user := &User{userid: 0, Username: username, x: 0, y: 0}
	for rows.Next() {
		return nil, errors.New("User has already exists.")
	}

	stmt, err := users.db.Prepare("insert into userinfo(username, password) values(?,?)")
	defer stmt.Close()
	if err != nil {
		return nil, err
	}

	_, err = stmt.Exec(username, password)
	if err != nil {
		return nil, err
	}

	return nil, nil
}
