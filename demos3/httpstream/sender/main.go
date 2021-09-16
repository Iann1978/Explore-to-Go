package main

import (
	"bytes"
	"fmt"
	"net/http"
)

func main() {
	c := http.Client{}
	data := []byte("This is a content that will be sent in the body")
	r, err := http.NewRequest("POST", "http://localhost:9090", bytes.NewBuffer(data))
	// You should never ignore the error returned by a call.
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	_, err = c.Do(r)
	if err != nil {
		fmt.Println(err)
	}

}
