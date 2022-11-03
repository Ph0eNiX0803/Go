package main

import (
	"fmt"
	"net"
)

func process(conn net.Conn) {
	defer conn.Close()
	for {
		buff := make([]byte, 1024)
		fmt.Println("服务器等待输入", conn.RemoteAddr())
		n, err := conn.Read(buff)
		if err != nil {
			fmt.Println("read err", err)
			return
		}
		fmt.Print(string(buff[:n]))
	}
}
func main() {
	fmt.Println("开始监听...")
	listen, err := net.Listen("tcp", "0.0.0.0:6666")
	if err != nil {
		fmt.Println("listen error")
		return
	}
	defer listen.Close()
	for {
		fmt.Println("等待连接")
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("err", err)
			//return
		} else {
			fmt.Printf("conn %v\n", conn.RemoteAddr())
		}
		go process(conn)
	}
	//fmt.Printf("listen = %v\n", listen)
}
