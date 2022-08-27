package myio

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"strings"
)

func Sample_ioreader() {
	conn, err := net.Dial("tcp", "rpcx.site:80")
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()
	var sb strings.Builder
	fmt.Fprintf(conn, "GET / HTTP 1.0\r\n\r\n")

	buf := make([]byte, 256)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			if err != io.EOF {
				log.Fatalln("error read", err)
			}
			break
		}
		sb.Write(buf[:n])
	}
	fmt.Println("response: ", sb.String())
	fmt.Println("total response size: ", sb.Len())
}

func Sample_iocopy() {
	conn, err := net.Dial("tcp", "rpcx.site:80")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Fprintf(conn, "GET / HTTP/1.1 \r\n\r\n")
	defer conn.Close()
	var sb strings.Builder
	// 可传递两个文件对象，实现文件的拷贝
	_, err = io.Copy(&sb, conn)
	if err != nil {
		log.Fatalln("read error: ", err)
	}
	fmt.Println("response: ", sb.String())
	fmt.Println("total response size: ", sb.Len())
}

func Sample_ioreadall() {
	conn, err := net.Dial("tcp", "rpcx.site:80")
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()
	fmt.Fprintf(conn, "GET / HTTP/1.1 \r\n\r\n")
	data, err := ioutil.ReadAll(conn)
	if err != nil {
		if err != io.EOF {
			log.Fatalln("read error: ", err)
		}
		panic(err)
	}

	fmt.Println("response: ", string(data))
	fmt.Println("total response size: ", len(data))
}

func Sample_ioreadfull() {
	conn, err := net.Dial("tcp", "rpcx.site:80")
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()
	fmt.Fprintf(conn, "GET / HTTP/1.1 \r\n\r\n")
	var sb strings.Builder
	buf := make([]byte, 256)
	for {
		n, err := io.ReadFull(conn, buf)
		if err != nil {
			if err != io.EOF && err != io.ErrUnexpectedEOF {
				log.Fatalln("error read ", err)
			}
			break
		}
		sb.Write(buf[:n])
	}
	fmt.Println("response: ", sb.String())
	fmt.Println("total response size: ", sb.Len())
}

func Sample_iomore() {
	conn, err := net.Dial("tcp", "rpcx.site:80")
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()
	fmt.Fprintf(conn, "GET / HTTP/1.1 \r\n\r\n")
	var sb strings.Builder
	buf := make([]byte, 256)
	rr := io.LimitReader(conn, 10000)
	for {
		n, err := io.ReadAtLeast(rr, buf, 256)
		if err != nil {
			if err != io.EOF && err != io.ErrUnexpectedEOF {
				log.Fatalln("error read ", err)
			}
			break
		}
		sb.Write(buf[:n])
	}
	fmt.Println("response: ", sb.String())
	fmt.Println("total response size: ", sb.Len())

}
