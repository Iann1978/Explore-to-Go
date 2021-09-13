package test

import (
	"errors"
	"fmt"
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
