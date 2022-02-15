package main

import (
	"fmt"
	"net"
	"strings"
)

func main() {
	serverMuti()
}

func serverMuti() {
	listener, err := net.Listen("tcp", "127.0.0.1:8088")
	if err != nil {
		fmt.Printf("服务器启动报错：%s", err)

		return
	}
	defer listener.Close()
	fmt.Println("服务器等待客户端连接！")

	for {

		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("监听出错", err)
			return
		}

		go handleConn(conn)
	}

}

func handleConn(conn net.Conn) {
	defer conn.Close()
	//获取连接客户端地址
	addr := conn.RemoteAddr()
	fmt.Println("client连接成功：", addr)
	buf := make([]byte, 4096)

	for {
		len, err := conn.Read(buf)
		if len == 0 {
			fmt.Println("客户端已经关闭！")
			return
		}
		if err != nil {
			fmt.Println("conn.read err", err)
			return
		}
		msg := string(buf[:len])
		fmt.Println("接收到到数据", msg)

		if strings.HasPrefix(msg, "exit") {
			fmt.Println("客户端请求退出！")
			return
		}

		upper := strings.ToUpper(msg)
		_, err = conn.Write([]byte(upper))
		if err != nil {
			fmt.Println("conn.Write err", err)
			return
		}

	}
}
func server() {

	fmt.Println("serve ready")
	listener, err := net.Listen("tcp", "127.0.0.1:8088")

	if err != nil {
		fmt.Printf("服务器启动报错：%s", err)
		return
	}
	fmt.Println("服务器等待客户端连接！")

	defer listener.Close()

	//阻塞监听连接请求
	conn, err := listener.Accept()
	if err != nil {
		fmt.Println("监听出错", err)
		return
	}

	fmt.Println("服务器建立连接成功！")
	defer conn.Close()

	//读取客户端发送数据
	buf := make([]byte, 4096)
	len, err := conn.Read(buf)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("服务器读取到数据", string(buf[:len]))

	conn.Write([]byte("hello client!"))
}
