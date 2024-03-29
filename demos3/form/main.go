package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Println(w, "Hello astasldk!")
}

func login(w http.ResponseWriter, r *http.Request) {
	
	fmt.Println("method:", r.Method)
	if r.Method == "GET" {
		t, _ := template.ParseFiles("login.gtpl")
		log.Println(t.Execute(w,nil))
	} else {
		err := r.ParseForm()
		if err != nil {
			fmt.Println("Error, Error, Error, Error!!!!!!!");
			log.Fatal("ParseForm: ", err)
			fmt.Fprintln(w,"False")
		} else {

			fmt.Println("username:", r.Form["username"])
			fmt.Println("password:", r.Form["password"])
			fmt.Fprintln(w, "Succeed")
		}

	}
}

func main() {
	http.HandleFunc("/", sayhelloName)
	http.HandleFunc("/login", login)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

