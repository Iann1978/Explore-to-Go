package test

import (
	"testing"
	"net/http"
)


func TestConnection(t *testing.T) {
	resp, err := http.Get("http://gobyexample.com")
	if err != nil {
		t.Errorf(err.Error())
	}
	defer resp.Body.Close()

}
