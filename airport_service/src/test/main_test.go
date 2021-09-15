package test

import (
	"airport_service/data"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"testing"
)

func getClientIp() (string, error) {
	addrs, err := net.InterfaceAddrs()

	if err != nil {
		return "", err
	}

	for _, address := range addrs {

		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String(), err
			}
		}
	}

	return "", errors.New("Can not find the client ip address!")
}

func TestConnection(t *testing.T) {
	ipstr, err := getClientIp()
	if err != nil {
		t.Errorf(err.Error())
	}

	req := fmt.Sprintf("http://%s:8099/hello", ipstr)
	fmt.Println(req)

	resp, err := http.Get(req)
	if err != nil {
		t.Errorf(err.Error())
	}
	defer resp.Body.Close()

}

func TestRebuildDb(t *testing.T) {
	ipstr, err := getClientIp()
	if err != nil {
		t.Errorf(err.Error())
	}

	req := fmt.Sprintf("http://%s:8099/rebuilddb", ipstr)
	fmt.Println(req)

	resp, err := http.Get(req)
	if err != nil {
		t.Errorf(err.Error())
	}
	defer resp.Body.Close()
}

type RegistResp struct {
	ErrorCode   data.ErrorCode
	ErrorString string
}

func TestRegist(t *testing.T) {

	aaa := &data.LoginResp2{"aaabbb"}
	fmt.Println(aaa)

	ipstr, err := getClientIp()
	if err != nil {
		t.Errorf(err.Error())
	}

	req := fmt.Sprintf("http://%s:8099/regist?username=test0&&password=test0", ipstr)
	fmt.Println(req)

	resp, err := http.Get(req)
	if err != nil {
		t.Errorf(err.Error())
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Errorf(err.Error())
	}

	fmt.Println("body:", string(body))

	var registResp RegistResp

	if err := json.Unmarshal(body, &registResp); err != nil {
		t.Errorf(err.Error())
		return
	}

	fmt.Println(registResp)

	if data.NoError != registResp.ErrorCode {
		t.Errorf(registResp.ErrorCode.String())
	}

}

func TestRegist_Case2(t *testing.T) {

	aaa := &data.LoginResp2{"aaabbb"}
	fmt.Println(aaa)

	ipstr, err := getClientIp()
	if err != nil {
		t.Errorf(err.Error())
	}

	req := fmt.Sprintf("http://%s:8099/regist?username=test0&&password=test0", ipstr)
	fmt.Println(req)

	resp, err := http.Get(req)
	if err != nil {
		t.Errorf(err.Error())
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Errorf(err.Error())
	}

	fmt.Println("body:", string(body))

	var registResp RegistResp

	if err := json.Unmarshal(body, &registResp); err != nil {
		t.Errorf(err.Error())
		return
	}

	fmt.Println(registResp)

	if data.UserExist != registResp.ErrorCode {
		t.Errorf(registResp.ErrorCode.String())
	}

}
