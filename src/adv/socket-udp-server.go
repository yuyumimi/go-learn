package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	updConcurrent()
}
func updConcurrent() {
	laddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:8003")
	if err != nil {
		fmt.Println("ResolveUDPAddr error:", err)
		return

	}

	conn, err := net.ListenUDP("udp", laddr)

	if err != nil {
		fmt.Println("ListenUDP", err)
		return
	}
	fmt.Println("服务器创建完成！")
	defer conn.Close()

	buf := make([]byte, 4096)
	for {

		//读取字节数，客户端地址
		len, addr, err := conn.ReadFromUDP(buf)
		if err != nil {
			fmt.Println("ReadFromUDP", err)
			return
		}

		fmt.Println("服务器读取%v %s", buf[:len], addr)

		now := time.Now().String()
		_, err = conn.WriteToUDP([]byte(now), addr)
		if err != nil {
			fmt.Println("WriteToUDP", err)
			return
		}
	}
}
func upd() {
	laddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:8003")
	if err != nil {
		fmt.Println("ResolveUDPAddr error:", err)
		return

	}

	conn, err := net.ListenUDP("udp", laddr)

	if err != nil {
		fmt.Println("ListenUDP", err)
		return
	}
	fmt.Println("服务器创建完成！")
	defer conn.Close()

	buf := make([]byte, 4096)
	//读取字节数，客户端地址
	len, addr, err := conn.ReadFromUDP(buf)
	if err != nil {
		fmt.Println("ReadFromUDP", err)
		return
	}

	fmt.Println("服务器读取%v %s", buf[:len], addr)

	now := time.Now().String()
	_, err = conn.WriteToUDP([]byte(now), addr)
	if err != nil {
		fmt.Println("WriteToUDP", err)
		return
	}
}
