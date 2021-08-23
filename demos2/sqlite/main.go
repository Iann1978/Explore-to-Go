//https://learnku.com/docs/build-web-application-with-golang/053-uses-the-sqlite-database/3183:x
package main

import ("database/sql"
	"fmt"
//	"time"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./foo.db")
	checkErr(err)

	// Insert the data
	stmt, err := db.Prepare("insert into userinfo(username, department, created) values(?,?,?)")
	checkErr(err)

	res, err := stmt.Exec("astaxie", "Developer Department", "2012-12-09")
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)

	fmt.Println(id)

	// Update the data
	stmt, err = db.Prepare("update userinfo set username=? where rowid=?")
	checkErr(err)

	res, err = stmt.Exec("astaxieupdate",id)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)

	// Query the data
	rows, err := db.Query("select rowid, username, department, time(created) from userinfo")
	checkErr(err)

	for rows.Next() {
		var rowid int
		var username string
		var department string
		//var created time.Time
		var tmp_created string
		err = rows.Scan(&rowid, &username, &department, &tmp_created)
		checkErr(err)
		fmt.Println(rowid)
		fmt.Println(username)
		fmt.Println(department)
		fmt.Println(tmp_created)
	}

	// Delete the data
	stmt, err = db.Prepare("delete from userinfo where rowid=?")
	checkErr(err)

	res, err = stmt.Exec(id)
	checkErr(err)

	affect, err = res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)

	db.Close()
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

