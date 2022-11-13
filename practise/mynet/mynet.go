package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
	"net/url"
	"os"
	"time"
)

func Sample_client(address string) {
	client_log := log.New(os.Stdout, "(client: )", log.LstdFlags|log.Lshortfile)
	// 连接到服务器
	client_log.Println("start connect to ", address)
	con, err := net.Dial("tcp", address)
	if err != nil {
		client_log.Fatal(err)
	}
	// con有一些方法Read, Write, Close, LocalAddr, RemoteAddr, SetDeadline
	con.SetDeadline(time.Now().Add(1 * time.Minute))

	client_log.Println("start to send message to server")
	// 向服务器发送信息
	fmt.Fprintf(con, "Hello server , I'am client\r\n\r\n")
	client_log.Println("start to get response from server")
	// 获取服务器发送过来的响应, 另外有con.Read(b[:])
	response, err := bufio.NewReader(con).ReadString('\n')
	if err != nil {
		client_log.Fatalf("error get response %v", err)
	}
	client_log.Printf("get response from server: %v\n", response)
	// 关闭连接
	con.Close()
	client_log.Println("end all")
}

func Sample_server(address string) {
	server_log := log.New(os.Stdout, "(server: )", log.Lshortfile|log.LstdFlags)
	// 开始监听
	server_log.Println("start to listen on ", address)
	ln, err := net.Listen("tcp", address)
	if err != nil {
		server_log.Fatalf("failed listen on %v : %v", address, err)
	}
	defer ln.Close()

	// 等待连接
	for {
		server_log.Println("wait client connect....")
		client_con, err := ln.Accept()
		if err != nil {
			server_log.Fatalf("failed accept from client: %v", err)
		}
		go func() {
			server_log.Println("client has connected, wait 20s to receive it's message")
			time.Sleep(time.Duration(20 * time.Second))
			// 获取客户端发送的请求
			client_message, err := bufio.NewReader(client_con).ReadString('\n')
			if err != nil {
				server_log.Fatalf("failed get message from client: %v", err)
			}
			server_log.Println("get message from client: ", client_message)
			// 向客户端发送响应, 另外可以用client_con.Write([]byte), io.WriteString(client_con, message)
			message := "Hello client, I have receive your message\r\n"
			server_log.Println("start to send message to client: ", message)
			fmt.Fprintf(client_con, message)
			// 关闭连接
			server_log.Println("start to close connection")
			client_con.Close()
		}()
	}
}

func Sample_utils() {
	// 组合主机和端口
	s := net.JoinHostPort("192.168.1.1", "8081")
	log.Println("(net.JoinHostPort) 192.168.1.1,8081 -> ", s)
	// 查询地址
	names, err := net.LookupAddr("127.0.0.1")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("(net.LookupAddr) 127.0.0.1 -> ", names)
	// 查询cname
	cname, err := net.LookupCNAME("127.0.0.1")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("(net.LookupCNAME) 127.0.0.1 -> ", cname)
	// 根据域名查询ip
	hosts, _ := net.LookupHost("www.baidu.com")
	log.Println("(net.LookupHost) www.baidu.com -> ", hosts)
	// 根据协议和服务名获取端口
	port, _ := net.LookupPort("tcp", "http")
	log.Println("(net.LookupPort) tcp,http -> ", port)
	// 获取主机的cidr
	ip, ipnet, _ := net.ParseCIDR("192.168.1.1/21")
	log.Println("(net.ParseCIDR) 192.168.1.1/21 -> ", ip, *ipnet)
	// 获取主机的mac
	mac, _ := net.ParseMAC("01:23:45:67:89:ab")
	log.Println("(net.ParseMAC) 01:23:45:67:89:ab -> ", mac)
}

func basic() {
	s := "http://www.baidu.com/aa/pp?ticket=cc"
	log.Printf("url: %v", s)
	result, err := url.Parse(s)
	if err != nil {
		log.Println(err)
	}
	log.Printf("Schema: %v", result.Scheme)
	log.Printf("Host: %v", result.Host)
	log.Printf("Path: %v", result.Path)
	log.Printf("RawQuery: %v", result.RawQuery)

}


func main() {
	var action string 
	var addr string 
	flag.StringVar(&action, "m", "", "设置运行方式，如server或client")
	flag.StringVar(&addr, "a", "", "设置运行地址")
	flag.Parse()
	if action == "" {
		flag.Usage()
		return 
	}
	switch action {
	case "server":
		Sample_server(addr)
	case "client":
		Sample_client(addr)
	default:
		fmt.Println("Error: -m shoud be server or action")
	}
}
