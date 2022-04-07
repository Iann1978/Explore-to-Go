package main

import (
	"net/http"
)

func main() {

	http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("./doc/"))))
	http.ListenAndServe(":81", nil)
}
