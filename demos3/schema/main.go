package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

	"github.com/gorilla/schema"
)

type Person struct {
	Username string
	Password string
	Age      int
}

func login(w http.ResponseWriter, r *http.Request) {

	var decoder = schema.NewDecoder()

	fmt.Println("method:", r.Method)
	if r.Method == "GET" {
		t, _ := template.ParseFiles("login.gtpl")
		log.Println(t.Execute(w, nil))
	} else {
		err := r.ParseForm()
		if err != nil {
			fmt.Println("Error, Error, Error, Error!!!!!!!")
			log.Fatal("ParseForm: ", err)
			fmt.Fprintln(w, "False")

		} else {

			//	fmt.Println("username:", r.Form["Username"])
			//	fmt.Println("password:", r.Form["Password"])
			//	fmt.Println("age:", r.Form["Age"])
			fmt.Fprintln(w, "Succeed")

			var person Person

			// r.PostForm is a map of our POST form values
			err = decoder.Decode(&person, r.PostForm)
			fmt.Println("username:", person.Username)
			fmt.Println("password:", person.Password)
			fmt.Println("age:", person.Age)
		}

	}
}

func main() {

	//	http.HandleFunc("/", sayhelloName)
	http.HandleFunc("/login", login)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
