package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("client ready")

	conn, err := net.Dial("udp", "127.0.0.1:8003")
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
