package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	recvFile()
}
func recvFile() {
	listen, err := net.Listen("tcp", "127.0.0.1:8008")
	if err != nil {
		fmt.Println("Listen", err)
		return
	}
	defer listen.Close()
	conn, err := listen.Accept()
	if err != nil {
		fmt.Println("Accept", err)
		return
	}

	defer conn.Close()
	fmt.Println("服务器启动成功！")
	buf := make([]byte, 4096)
	len, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Read", err)
		return
	}
	conn.Write([]byte("ok"))

	bytes := buf[:len]
	filename := string(bytes)

	saveFile(conn, filename)
}

func saveFile(conn net.Conn, filename string) {
	file, err := os.Create("f:\\tmp\\" + filename)
	if err != nil {
		fmt.Println("Create ", err)
		return
	}

	buf := make([]byte, 4096)
	for {

		len, err := conn.Read(buf)
		if err != nil {
			if len == 0 {
				fmt.Println("接收文件完成 ")

				return
			}
			fmt.Println(" conn Read ", err)
			return
		}

		_, err = file.Write(buf[:len])
		if err != nil {
			fmt.Println("Write ", err)
			return
		}
	}
}
