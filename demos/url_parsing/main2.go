package main

import (
	"fmt"
//	"net"
	"net/url"
)

func main() {
	
	s := "http://172.31.231.247:8099/login?key=123&username=aaa&&password=bbb"
	fmt.Println("url:", s)

	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}
	fmt.Println("Url:", u.Scheme)

	fmt.Println("Query:", u.RawQuery)

	m, _ := url.ParseQuery(u.RawQuery)
	fmt.Println("ParsedQuery:", m)

	fmt.Println("key:", m["key"])
	fmt.Println("username:", m["username"])
	fmt.Println("password:", m["password"])


	for k, v := range m {
		fmt.Println("k:", k, ", v:", v[0])
	}



}
