package main

import (
	"fmt"
	"net"
	"os"
	//"io"
)

func main() {
	fmt.Println("Start Go's Tcp Server.")
	var l net.Listener;
	var err error
	l, err = net.Listen("tcp", ":1200")
	if err != nil {
		fmt.Println("Error listening:", err)
		os.Exit(1)
	}
	var conn net.Conn
	for {
		conn, err = l.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err)
			os.Exit(1)
		}
		fmt.Printf("Received message %s -> %s \n", conn.RemoteAddr(), conn.LocalAddr())
		handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	defer conn.Close()
	fmt.Println("handleRequest.\n")
	buffer := make([]byte,1024)
	for {
		msg, err := conn.Read(buffer)

		if err != nil {
			fmt.Println("connection err!:", err)
			return
		}
		fmt.Print(string(buffer[:msg]))
		//io.Copy(conn,conn)
	}

}
