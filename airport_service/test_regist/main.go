package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"sync"
	"time"
)

func RegistProcess(wg *sync.WaitGroup, base int, count int) {
	defer wg.Done()

	for i := 0; i < count; i++ {
		username := "user_" + strconv.FormatInt(int64(base+i), 10)
		resp, err := http.PostForm("http://192.168.100.80:8099/regist", url.Values{
			"username":  {username},
			"password":  {"user"},
			"time":      {"88224646"},
			"usergroup": {"1"},
		})

		if err != nil {
			fmt.Println(err.Error())
			return
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		fmt.Println("body:", string(body))

	}

}

func main() {
	fmt.Println("hehe")

	start := time.Now()

	var wg sync.WaitGroup

	count := 1000
	cocount := 10
	total := cocount * count

	for i := 0; i < cocount; i++ {
		base := 10000 * i

		wg.Add(1)
		go RegistProcess(&wg, base, count)

	}

	wg.Wait()
	end := time.Now()
	duation := end.Sub(start)
	tps := float64(total) / duation.Seconds()
	fmt.Println("count:", total)
	fmt.Println("duation:", duation.Seconds())
	fmt.Println("tps:", tps)

}
