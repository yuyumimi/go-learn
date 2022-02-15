package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	startClientMuti()
}

func startClientMuti() {
	fmt.Println("client ready")
	conn, err := net.Dial("tcp", "127.0.0.1:8088")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("连接服务器成功！")
	defer conn.Close()

	go func() {
		buf := make([]byte, 4096)
		for {
			n, err := os.Stdin.Read(buf)
			if err != nil {
				fmt.Println("os.Stdin.Read err", err)
				continue
			}
			conn.Write(buf[:n])
		}
	}()

	buf := make([]byte, 4096)
	for {
		n, err := conn.Read(buf)
		if n == 0 {
			fmt.Println("服务器端关闭！")
			return
		}
		if err != nil {
			fmt.Println("conn.Read err", err)
			return
		}
		fmt.Println("服务器返回数据：", string(buf[:n]))
	}
}

func startClient() {

	fmt.Println("client ready")

	conn, err := net.Dial("tcp", "127.0.0.1:8088")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("客户端连接成功！")
	defer conn.Close()

	//发送数据给服务器
	conn.Write([]byte("hello server!"))

	//接收服务器数据
	buf := make([]byte, 4096)
	len, err := conn.Read(buf)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("接收到服务器数据：", string(buf[:len]))
}
