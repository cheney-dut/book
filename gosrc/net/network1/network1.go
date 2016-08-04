// network1
// 根据IP和掩码获得网络
package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	name := "192.168.1.97"
	ip := net.ParseIP(name)
	mask := ip.DefaultMask()
	network := ip.Mask(mask)

	// 192.168.1.0
	fmt.Fprintf(os.Stdout, "network: %s\n", network.String())

	// ip:      192.168.1.97
	// mask:    255.255.255.0
	// network: 192.168.1.0
}
