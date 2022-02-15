package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	downloadFile()
}
func sendFile(conn net.Conn, filename string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Open ", err)
		return
	}
	defer file.Close()

	buf := make([]byte, 4096)
	for {

		len, err := file.Read(buf)
		if err != nil {
			if err == io.EOF {
				fmt.Println("发送文件完成 ")

				return
			}
			fmt.Println("Read ", err)
			return
		}

		_, err = conn.Write(buf[:len])
		if err != nil {
			fmt.Println("Write ", err)
			return
		}
	}
}

func downloadFile() {
	args := os.Args

	if len(args) != 2 {
		fmt.Println("格式为：go run ***.go 文件名称")
		return
	}

	path := args[1]
	fmt.Println("发送的文件：", path)

	fileInfo, err := os.Stat(path)
	if len(args) != 2 {
		fmt.Println("Stat :", err)
		return
	}
	name := fileInfo.Name()

	conn, err := net.Dial("tcp", "127.0.0.1:8008")
	if err != nil {
		fmt.Println("Dial", err)
		return
	}

	defer conn.Close()
	fmt.Println("连接服务器成功")

	conn.Write([]byte(name))

	buf := make([]byte, 16)
	len, err := conn.Read(buf)
	ok := string(buf[:len])
	fmt.Println("服务器返回：", ok)

	if ok == "ok" {
		sendFile(conn, path)
	}

}
