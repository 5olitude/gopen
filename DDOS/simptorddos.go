package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"sync"
	"sync/atomic"
	"time"

	"golang.org/x/net/proxy"
)

var wg sync.WaitGroup
var req uint64

func main() {
	tbProxyURL, _ := url.Parse("socks5://127.0.0.1:9050")
	tbDialer, _ := proxy.FromURL(tbProxyURL, proxy.Direct)
	tbTransport := &http.Transport{Dial: tbDialer.Dial}
	client := &http.Client{Transport: tbTransport}
	for i := 1; i <= 200; i++ {
		go func() {
			for {
				resp, _ := client.Get("gggggg")
				n, _ := ioutil.ReadAll(resp.Body)
				fmt.Println(string(n))
				resp.Body.Close()
				atomic.AddUint64(&req, 1)
				fmt.Println(req)
			}
		}()
		time.Sleep(time.Second)
	}
}
