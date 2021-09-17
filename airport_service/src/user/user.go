package user

import (
	"airport_service/data"
	"crypto/rand"
	"database/sql"
	"encoding/base64"
	"errors"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	userid   int32
	Username string
	Session  string
	x        float64
	y        float64
}

type UserDatabase struct {
	db          *sql.DB
	onlineUsers map[int]*User
	counter     int32
}

type UserSetError struct {
	ErrorCode   data.ErrorCode
	ErrorString string
}

func (e *UserSetError) Error() string {
	return e.ErrorString
}

type UserSet interface {
	Regist(username string, password string) (*User, error)
	UserLongin(username string, password string) (*User, error)
	UserLogout(username string) error
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

// GenerateRandomBytes returns securely generated random bytes.
// It will return an error if the system's secure random
// number generator fails to function correctly, in which
// case the caller should not continue.
func GenerateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	// Note that err == nil only if we read len(b) bytes.
	if err != nil {
		return nil, err
	}

	return b, nil
}

// GenerateRandomString returns a URL-safe, base64 encoded
// securely generated random string.
// It will return an error if the system's secure random
// number generator fails to function correctly, in which
// case the caller should not continue.
func GenerateRandomString(s int) (string, error) {
	b, err := GenerateRandomBytes(s)
	return base64.URLEncoding.EncodeToString(b), err
}
func (users *UserDatabase) UserLongin(username string, userPassword string) (*User, error) {

	rows, err := users.db.Query("select username, password from userinfo where username=?", username)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	session, _ := GenerateRandomString(32)

	user := &User{userid: 0, Username: username, x: 0, y: 0, Session: session}

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

				break

				return user, nil
			} else {
				return nil, errors.New("Password error!!!")
			}

		}
	}

	rows.Close()
	stmt, err := users.db.Prepare("update userinfo set session=? where username=?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(session, username)
	if err != nil {
		return nil, err
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
		e := UserSetError{ErrorCode: data.UserExist, ErrorString: data.UserExist.String()}
		return nil, &e
		//return nil, errors.New("User has already exists.")
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

func (users *UserDatabase) UserLogout(reqUsername string) error {

	var username string
	var session *string
	row := users.db.QueryRow("select username, session from userinfo where username=?", reqUsername)

	err := row.Scan(&username, &session)
	if err != nil {
		fmt.Println(err)
		e := UserSetError{ErrorCode: data.UnknownError, ErrorString: data.UnknownError.String()}
		return &e
	}

	if session == nil || len(*session) == 0 {
		e := UserSetError{ErrorCode: data.UserOffline, ErrorString: data.UserOffline.String()}
		return &e
	}

	stmt, err := users.db.Prepare("update userinfo set session=null where username=?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(username)
	if err != nil {
		return err
	}

	return nil
	//  err {
	// case sql.ErrNoRows:
	//   fmt.Println("No rows were returned!")
	// case nil:
	//   fmt.Println(id, email)
	// default:
	//   panic(err)

	// if err != nil {
	// 	fmt.Println(err)
	// 	e := UserSetError{ErrorCode: data.UnknownError, ErrorString: data.UnknownError.String()}
	// 	return &e
	// }

}
