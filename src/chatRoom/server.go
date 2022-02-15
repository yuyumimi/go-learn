package main

import (
	"fmt"
	"net"
	"runtime"
	"strings"
	"time"
)

type Client struct {
	C    chan string
	Name string
	Addr string
}

var onlineMap map[string]Client

var message = make(chan string)

func main() {
	listener, err := net.Listen("tcp", "127.0.0.1:8088")
	if err != nil {
		fmt.Printf("服务器启动报错：%s", err)

		return
	}
	defer listener.Close()
	fmt.Println("服务器等待客户端连接！")

	go manager()
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener.Accept err:", err)
			return
		}

		go handleConnect(conn)
	}
}

func handleConnect(conn net.Conn) {
	defer conn.Close()

	addr := conn.RemoteAddr().String()

	client := Client{make(chan string), addr, addr}
	fmt.Printf("客户端%s连接\n", addr)

	onlineMap[addr] = client
	go listenWChannel(client, conn)

	message <- makeMsg(client, "login")

	quit := make(chan bool)
	active := make(chan bool)

	go func() {
		buf := make([]byte, 4096)
		for {

			l, err := conn.Read(buf)
			if l == 0 {
				quit <- true
				fmt.Printf("客户端%s退出\n", client.Name)
				return
			}
			if err != nil {

				fmt.Println("client read err:", err)
				return
			}
			msg := string(buf[:l])
			msg = strings.ReplaceAll(msg, "\r\n", "")
			if msg == "who" {
				conn.Write([]byte("user list: \n"))
				for _, val := range onlineMap {
					user := val.Name
					conn.Write([]byte(user + " \n"))

				}

			} else if strings.HasPrefix(msg, "rename|") {
				split := strings.Split(msg, "|")
				if len(split) == 2 {
					client.Name = split[1]
					onlineMap[client.Addr] = client
					conn.Write([]byte("rename ok!"))
				}
			} else {
				message <- makeMsg(client, msg)

			}
			active <- true
		}
	}()
	for {

		select {
		case <-quit:
			close(client.C)
			delete(onlineMap, client.Addr)
			message <- makeMsg(client, "logout")
			runtime.Goexit()
		case <-active:
			break
		case <-time.After(time.Second * 60):
			close(client.C)
			delete(onlineMap, client.Addr)
			message <- makeMsg(client, "logout")
			runtime.Goexit()
		}
	}
}

func makeMsg(client Client, msg string) (buf string) {

	buf = "[" + client.Addr + "]" + client.Name + ":" + msg + "\n"
	return
}

func listenWChannel(c Client, conn net.Conn) {
	for msg := range c.C {
		conn.Write([]byte(msg + "\n"))
	}
}

func manager() {
	onlineMap = make(map[string]Client)

	for {

		msg := <-message

		for _, client := range onlineMap {
			ch := client.C

			ch <- msg
		}
	}

}
