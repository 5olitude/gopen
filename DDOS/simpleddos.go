package main

import (
	"fmt"
	"net/http"
	"sync"
	"sync/atomic"
	"time"
)

var wg sync.WaitGroup
var req uint64

func main() {
	for i := 1; i <= 100; i++ {
		go func() {
			for {
				_, _ = http.Get("gfgfgfgfgf")
				atomic.AddUint64(&req, 1)
				fmt.Println(req)
			}
		}()
		time.Sleep(1 * time.Second)
	}
}
