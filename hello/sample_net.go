package main

import (
	"fmt"
	"net"
	"bufio"
)

func listen(){
	ln, err := net.listen("tcp", ":8080")
	if err != nil {
		fmt.Println("failed establish listen")
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("accept connection error")
		}
		result := bufio.NewReader(conn).ReadString('\n')
		fmt.Println("get info from conn: ", result)
	}
}

func connect(){
	conn, err := net.Dial("tcp", "www.baidu.com:80")
	if err != nil {
		fmt.Printf("connect failed")
	}
	fmt.Fprintf(conn, "GET / HTTP/1.0\r\n\r\n")
	status, err := bufio.NewReader(conn).ReadString('\n')
	if err != nil {
		fmt.Printf("get string from conn failed")
	}
	fmt.Print("result: ", status)
}

func main() {
	fmt.Println("start to listen on :80....")
}
