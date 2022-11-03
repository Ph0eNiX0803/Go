package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "114.213.210.10:6666")
	if err != nil {
		fmt.Println("client dial err", err)
		return
	}
	fmt.Println("conn success", conn.RemoteAddr())
	for {
		reader := bufio.NewReader(os.Stdin)

		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("read err", err)
		}
		if line == "exit\r\n" {
			return
		}
		n, err := conn.Write([]byte(line))
		if err != nil {
			fmt.Println("conn.write err", err)
		}
		fmt.Printf("发送了 %d 字节数据\n", n)
	}

}
