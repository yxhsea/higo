package main

import (
	"log"
	"math/rand"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func NewMultipleHostsReversProxy(targets []*url.URL) *httputil.ReverseProxy {
	director := func(req *http.Request) {
		target := targets[rand.Int()%len(targets)]
		req.URL.Scheme = target.Scheme
		req.URL.Host = target.Host
	}

	return &httputil.ReverseProxy{Director: director}
}

func main() {
	proxy := NewMultipleHostsReversProxy([]*url.URL{
		{
			Scheme: "http",
			Host:   "127.0.0.1:8081",
		}, {
			Scheme: "http",
			Host:   "127.0.0.1:8082",
		},
	})

	http.Handle("/abc", proxy)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
