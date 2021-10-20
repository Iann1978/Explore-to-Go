package main

import (
	"net/http"
	"net/url"
	"fmt"
	"bufio"
)

func main() {

	resp, err := http.PostForm("http://127.0.0.1:9090/login", url.Values{"Username":{"aaa"}, "Password":{"password"}, "Age":{"23abc"}})
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	scanner := bufio.NewScanner(resp.Body)
	for i := 0; scanner.Scan() && i<5; i++ {
		fmt.Println(scanner.Text())
	}

}
