package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func startReverseProxy(config TestGenieConfig) {
	origin, _ := url.Parse(config.Host.URL)
	director := func(req *http.Request) {
		req.URL.Scheme = "http"
		req.URL.Host = origin.Host
		req.Host = origin.Host
	}

	fmt.Printf("Routing to host: %v\n", origin.Host)

	proxy := &httputil.ReverseProxy{Director: director}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if config.Host.Live {
			_directory = r.URL.Path
			proxy.ServeHTTP(NewCustomWriter(w), r)
		} else {
			panic("Feature under development!")
		}
	})
}
