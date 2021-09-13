package main

import (
	"airport_service/mydb"
	"fmt"
	"net/http"
)

func rebuilddb(w http.ResponseWriter, req *http.Request) {

	fmt.Fprintln(w, "hehe")
	mydb.RemoveDb()

	mydb.CreateDb()
}
