package main

import (
	"fmt"
	"net"
	"os"
)

func main() {

	conn, err := net.Dial("tcp", "127.0.0.1:8088")
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	go func() {
		buf := make([]byte, 1024)
		for {
			n, err := conn.Read(buf)
			if n == 0 {
				fmt.Println("服务器端退出！")
				return
			}
			if err != nil {
				fmt.Println("conn.Read error:", err)
				return
			}

			_, err = os.Stdout.Write(buf[:n])
			if err != nil {
				fmt.Println("Stdout.Read error:", err)
				return
			}

		}
	}()

	fmt.Printf("请输入")
	stdin := os.Stdin

	buf := make([]byte, 2048)

	for {

		len, err := stdin.Read(buf)

		if err != nil {
			fmt.Println("stdin.Read error:", err)
			return
		}

		input := buf[:len]

		_, err = conn.Write(input)
		if err != nil {
			fmt.Println("conn.Write error:", err)
			return
		}
	}
}
