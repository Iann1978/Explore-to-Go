package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/golang/protobuf/proto"
)

func main() {
	resp, err := http.Get("http://localhost:8080/login")
	if err != nil {
		fmt.Println(err)
	} else {
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(string(body))
		user := &Student{}
		proto.UnmarshalMerge(body, user)
		fmt.Println("name:", user.Name)
		fmt.Println("male:", user.Male)
		// else {
		// 	user := &module.User{}
		// 	proto.UnmarshalMerge(body, user)
		// 	fmt.Println(*user)
		// }

	}
}
