//Basic TOR Example
package main

import (
	"bytes"
	"log"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/proxy"
)

func main() {
	// Instantiate default collector
	c := colly.NewCollector(colly.AllowURLRevisit())

	// Rotate two socks5 proxies
	rp, err := proxy.RoundRobinProxySwitcher("socks5://127.0.0.1:9050", "socks5://127.0.0.1:9050")
	if err != nil {
		log.Fatal(err)
	}
	c.SetProxyFunc(rp)

	// Print the response
	c.OnResponse(func(r *colly.Response) {
		log.Printf("%s\n", bytes.Replace(r.Body, []byte("\n"), nil, -1))
	})

	for i := 0; i < 1; i++ {
		c.Visit("https://httpbin.org/ip")
	}
}
