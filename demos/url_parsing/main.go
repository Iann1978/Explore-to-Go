package main

import (
	"fmt"
	"net"
	"net/url"
)

func main() {

	s := "postgred://user:pass@host.com:5432/path?k=v#f"

	u, err := url.Parse(s)
	if err != nil {
		panic(err)
	}

	fmt.Println("Url:", s)
	fmt.Println("Scheme:", u.Scheme)
	fmt.Println("User:", u.User)
	fmt.Println("User.Username:", u.User.Username())
	p, _ := u.User.Password()
	fmt.Println("User.Password:", p)

	fmt.Println("Host:", u.Host)
	host, port, _ := net.SplitHostPort(u.Host)
	fmt.Println("Host.host:", host)
	fmt.Println("Host.port:", port)

	fmt.Println("Path:", u.Path)
	fmt.Println("Fragment:", u.Fragment)
	
	fmt.Println("RawQuery:", u.RawQuery)
	m, _ := url.ParseQuery(u.RawQuery)
	fmt.Println("ParsedQuery:", m)
	fmt.Println("m[k][0]:", m["k"][0])
}
