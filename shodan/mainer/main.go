package main

import (
	"fmt"
	"log"
	"os"
	"github.com/josephnedher/blackhatgo/httpgo/shodan"
)
func main(){
	if len(os.Args)!=2{
		log.Fatalln("usage : main <seach term>")
	}
	apiKey:=os.Getenv("shodan_api_key")
	s:=shodan.New(apiKey)
	info,err:=s.APIInfo()
	if err != nil{
		log.Panicln
	}
	fmt.Printf(
		"Query Cedits :%d \n scan Credits: %d \n\n",info.QueryCredits,info.ScanCredits)
	)
	hostsearch,err:=s.HostSearch(os.Args[1])
	if err != nil {
		log.Panicln(err)
	}
	for _,host:range hostsearch.Matches{
		fmt.Printf("%18s%8d\n", host.IPString, host.Port)

	}

}