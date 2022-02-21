//https://learnku.com/docs/build-web-application-with-golang/053-uses-the-sqlite-database/3183:x
package main

import (
	"database/sql"
	"fmt"
	"os"
	"sync"
	"time"

	//	"time"
	_ "github.com/mattn/go-sqlite3"
)

func resetdb() *sql.DB {
	os.Remove("foo.db")

	db, err := sql.Open("sqlite3", "./foo.db")
	checkErr(err)

	sqlstr := `create table userinfo(
		userid integer primary key autoincrement,
		username text,
		password text, 
		session text,
		longitude real default 0.0,
		latitude real default 0.0,
		time integer default 0,
		speed integer default 0,
		orientation integer default 0,
		task_id integer default 0,
		time_remain int32  default 0,
		task_level  int32  default 0,

		msg text,
		status integer  default 0)`
	stmt, err := db.Prepare(sqlstr)
	defer stmt.Close()
	checkErr(err)

	_, err = stmt.Exec()
	checkErr(err)
	return db
}

type ProfileResult struct {
	test_count int32
	start      int64
	end        int64
	tps        int32
}

func insert_test1(db *sql.DB) *ProfileResult {
	start := time.Now().Unix()
	fmt.Println(start)

	test_count := 1000
	for i := 0; i < test_count; i++ {
		stmt, err := db.Prepare("insert into userinfo(username, password) values(?,?)")
		if err != nil {
			fmt.Println(err)
			panic(err)
		}
		_, err = stmt.Exec("aaa", "bbb")
		if err != nil {
			fmt.Println(err)
			panic(err)
		}
		stmt.Close()

	}
	end := time.Now().Unix()
	fmt.Println(end)

	fmt.Println("tps:", float32(test_count)/(float32(end-start)))

	result := ProfileResult{
		test_count: int32(test_count),
		start:      start,
		end:        end,
		tps:        int32(float32(test_count) / (float32(end - start))),
	}

	return &result

}

func insert_test2(db *sql.DB) *ProfileResult {
	start := time.Now().Unix()
	fmt.Println(start)

	test_count := 1000
	stmt, err := db.Prepare("insert into userinfo(username, password) values(?,?)")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	for i := 0; i < test_count; i++ {

		_, err = stmt.Exec("aaa", "bbb")
		if err != nil {
			fmt.Println(err)
			panic(err)
		}
	}
	stmt.Close()
	end := time.Now().Unix()
	fmt.Println(end)

	fmt.Println("tps:", float32(test_count)/(float32(end-start)))

	result := ProfileResult{
		test_count: int32(test_count),
		start:      start,
		end:        end,
		tps:        int32(float32(test_count) / (float32(end - start))),
	}

	return &result

}

func insert_test3(db *sql.DB) *ProfileResult {
	start := time.Now().Unix()
	fmt.Println(start)
	mutex := sync.Mutex{}
	co_count := 100
	test_count := 10000
	test_count_per_co := test_count / co_count
	var wg sync.WaitGroup
	for i := 0; i < co_count; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			stmt, err := db.Prepare("insert into userinfo(username, password) values(?,?)")
			if err != nil {
				fmt.Println(err)
				panic(err)
			}

			for i := 0; i < test_count_per_co; i++ {
				mutex.Lock()
				_, err = stmt.Exec("aaa", "bbb")
				mutex.Unlock()
				if err != nil {
					fmt.Println(err)
					panic(err)
				}
			}
			stmt.Close()

		}()

	}
	wg.Wait()

	end := time.Now().Unix()
	fmt.Println(end)

	fmt.Println("tps:", float32(test_count)/(float32(end-start)))

	result := ProfileResult{
		test_count: int32(test_count),
		start:      start,
		end:        end,
		tps:        int32(float32(test_count) / (float32(end - start))),
	}

	return &result

}

func main() {
	db := resetdb()
	result := insert_test1(db)
	fmt.Println("insert_test1:", result)

	db = resetdb()
	result = insert_test2(db)
	fmt.Println("insert_test2:", result)

	db = resetdb()
	result = insert_test3(db)
	fmt.Println("insert_test3:", result)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
