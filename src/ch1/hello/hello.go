package main

import (
	"net/http"
	_ "net/http/pprof"
	"time"
)

func main() {
	ticker := time.NewTicker(time.Second * 20)
	defer ticker.Stop()
	go func() {
		for range ticker.C {
			go discover()
		}
	}()
	http.ListenAndServe("localhost:8888", nil)
}

func discover() {
	ch := make(chan int, 1)
	ch <- 1
}
