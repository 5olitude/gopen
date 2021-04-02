package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"golang.org/x/net/proxy"
)

func main() {
	urlpse, _ := url.Parse("socks5://127.0.0.1:9050")
	wayer, _ := proxy.FromURL(urlpse, proxy.Direct)
	nwpath := &http.Transport{Dial: wayer.Dial}
	client := &http.Client{Transport: nwpath}
	resp, _ := client.Get("https://httpbin.org/ip")
	respon, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(respon))
}
