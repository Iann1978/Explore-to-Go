package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func receivestream(w http.ResponseWriter, r *http.Request) {

	d, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	fmt.Println("Request content : ", string(d))
}

func main() {
	fmt.Println("receive server start.")
	http.HandleFunc("/", receivestream)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
