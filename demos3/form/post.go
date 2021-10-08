package main

import (
	"net/http"
	"net/url"
	"fmt"
	"bufio"
)

func main() {

	resp, err := http.PostForm("http://127.0.0.1:9090/login", url.Values{"username":{"aaa"}, "password":{"password"}})
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	scanner := bufio.NewScanner(resp.Body)
	for i := 0; scanner.Scan() && i<5; i++ {
		fmt.Println(scanner.Text())
	}

}
