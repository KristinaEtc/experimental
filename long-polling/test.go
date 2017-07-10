package main

import (
	"flag"
	"fmt"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"
)

var (
	nurl *url.URL
	u    = flag.String("url", "http://127.0.0.1:7000/", "Connection URL")
	c    = flag.Int("c", 10000, "Connection Count")
)

// Connect create connect and process it
func Connect() {
	req := http.Request{
		Method: "GET",
		Header: http.Header{},
		URL:    nurl,
	}
	for {
		dial, err := net.Dial("tcp", nurl.Host)
		if err != nil {
			fmt.Println(err)
			return
		}
		conn := httputil.NewClientConn(dial, nil)
		conn.Do(&req)
		conn.Close()
	}
}

func main() {
	flag.Parse()
	nurl, _ = url.Parse(*u)
	for i := 0; i < *c; i++ {
		go Connect()
	}
	time.Sleep(1e10)
}
