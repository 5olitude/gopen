//This example is taken from the book "Black Hat Go"
//non-concurrent scanner
package main

import (
	"fmt"
	"net"
)

func main() {
	for i := 1; i <= 1024; i++ {
		address := fmt.Sprintf("scanme.nmap.org:%d", i)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			continue
		}
		conn.CLose()
		fmt.Printf("%d open\n", i)
	}
}
