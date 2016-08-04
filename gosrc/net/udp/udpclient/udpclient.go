// udpclient
package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	msg := "你好啊！！！"
	if len(os.Args) > 1 {
		msg = os.Args[1]
	}
	fmt.Println("msg:", msg)

	pUDPAddr, err := net.ResolveUDPAddr("udp", "127.0.0.1:7070")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Erros ResolveUDPAddr: %s\n", err.Error())
		return
	}
	pUDPConn, err := net.DialUDP("udp", nil, pUDPAddr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error DialUDP: %s\n", err.Error())
		return
	}
	defer pUDPConn.Close()

	n, err := pUDPConn.Write([]byte(msg))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error WriteToUDP: %s", err.Error())
		return
	}
	fmt.Fprintf(os.Stdout, "writed: %d\n", n)
	buf := make([]byte, 1024)
	n, _, err = pUDPConn.ReadFromUDP(buf)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error ReadFromUDP: %s\n", err.Error())
		return
	}
	fmt.Fprintf(os.Stdout, "readed: %d %s\n", n, string(buf[:n]))
}
