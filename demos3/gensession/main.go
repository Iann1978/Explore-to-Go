package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
)

func main() {
	fmt.Println("Hello, session id")

	token, err := GenerateRandomString(32)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(token)
}

func GenerateRandomBytes(n int) ([]byte, error) {

	b := make([]byte, n)
	_, err := rand.Read(b)

	if err != nil {
		return nil, err
	}

	return b, nil
}

func GenerateRandomString(s int) (string, error) {
	b, err := GenerateRandomBytes(s)
	return base64.URLEncoding.EncodeToString(b), err
}
