package main

import (
	"fmt"
	"net"
	"net/rpc"
)

/*
服务端使用步骤
1.注册rpc服务对象，给对象绑定方法（定义类，绑定类方法）
2.创建监听
3.建立连接
4.将连接绑定rpc服务

客户端使用步骤
1.用rpc连接服务器
2.调用远程函数

*/

type Human struct {
}

func (this *Human) getName(name string, resp *string) error {

}

func main() {
	err := rpc.RegisterName("rpc-test", new(Human))

	if err != nil {
		fmt.Println("regis err", err)
		return
	}
	listen, err := net.Listen("tcp", "127.0.0.1")
	if err != nil {
		fmt.Println("Listen err", err)
		return
	}
	defer listen.Close()
	conn, err := listen.Accept()
	if err != nil {
		fmt.Println("Accept err", err)
		return
	}
	defer conn.Close()
	rpc.ServeConn(conn)

}
