//Included only  thr free plan api of ipnfo.io

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
)

type IpInfo struct {
	Ip          string `json:"ip"`
	HostName    string `json:"hostname"`
	AnyCast     string `json:"anycast"`
	City        string `json:"city"`
	Region      string `json:"region"`
	Country     string `json:"country"`
	Location    string `json:"loc"`
	Orgnisation string `json:"org"`
	Postal      string `json:"postal"`
	TimeZone    string `json:"timezone"`
}

func main() {
	Ipaddr := flag.String("ip", "", "Enter the Ip ")
	token := flag.String("token", "", "Enter the token")
	flag.Parse()
	resp, err := http.Get("https://ipinfo.io/" + *Ipaddr + "/json?token=" + *token)
	if err != nil {
		fmt.Println("Something Bad occured oops")
	}
	var Extract IpInfo
	if err := json.NewDecoder(resp.Body).Decode(&Extract); err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Ip:		%s\n", Extract.Ip)
	fmt.Printf("HostName:	%s\n", Extract.HostName)
	fmt.Printf("AnyCast:	%s\n", Extract.AnyCast)
	fmt.Printf("CITY:		%s\n", Extract.City)
	fmt.Printf("REGION:		%s\n", Extract.Region)
	fmt.Printf("COUNTRY:	%s\n", Extract.Country)
	fmt.Printf("LOCAION:	%s\n", Extract.Location)
	fmt.Printf("ORGANISATION:	%s\n", Extract.Orgnisation)
	fmt.Printf("POSTAL:		%s\n", Extract.Postal)
	fmt.Printf("TIMEZONE:	%s\n", Extract.TimeZone)

}

// Usage  go run ipinfofree.go -ip=127.0.0.1 -token=.klkrktjjrht
