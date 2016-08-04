// network3
package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	// 查看主机服务器（service）占用的端口,这些服务，都是tcp或者udp的
	port, err := net.LookupPort("tcp", "telnet") // 查看telnet服务器使用的端口

	if err != nil {
		fmt.Fprintf(os.Stderr, "未找到指定服务")
		return
	}

	fmt.Fprintf(os.Stdout, "telnet port: %d \n", port)
	fmt.Println("-------------------------------------")

	// 将一个host地址转换为TCPAddr。host=ip:port
	pTCPAddr, err := net.ResolveTCPAddr("tcp", "www.baidu.com:80")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Err: %s", err.Error())
		return
	}
	fmt.Fprintf(os.Stdout, "www.baidu.com:80 IP: %s PORT: %d \n", pTCPAddr.IP.String(), pTCPAddr.Port)
}
