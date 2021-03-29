// Beautifullsoup of golang colly
//The example is an automated login with tor
//Ensure you  downloaded the tor service
// start the tor service using >>> service tor restart
// check the tor status by >>>>    serive tor  status
// To check the service started properly you should check the traffic or run the basic tor request demo torswe.go i added the file

package main

import (
	"fmt"
	"log"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/proxy"
)

func main() {
	//initiates the configuration
	c := colly.NewCollector(colly.AllowURLRevisit())
	//defining the proxy chain
	revpro, err := proxy.RoundRobinProxySwitcher("socks5://127.0.0.1:9050", "socks5://127.0.0.1:9050")
	if err != nil {
		log.Fatal(err)
	}
	c.SetProxyFunc(revpro)
	//parsing the required field from html we are extracting the csrf_token required for the login
	c.OnHTML("form[role=form] input[type=hidden][name=CSRF_TOKEN]", func(e *colly.HTMLElement) {
		csrftok := e.Attr("value")
		fmt.Println(csrftok)
		//posting the csrf value along with password
		err := c.Post("https://www.something.com/login.jsp", map[string]string{"CSRF_TOKEN": csrftok, "username": "username", "password": "password"})
		if err != nil {
			log.Fatal(err)
		}
		return
	})
	//The website to visit
	c.Visit("https://www.something.com/login.jsp")
	//maintaining the connection using clone not initiating a callback request
	d := c.Clone()
	d.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		fmt.Printf("Link found: %q -> %s\n", e.Text, link)

	})

	d.Visit("https://skkskskskk.htm")
}
