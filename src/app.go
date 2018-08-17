package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {
	config := getConfig()
	origin, _ := url.Parse(config.Host.URL)
	director := func(req *http.Request) {
		req.URL.Scheme = "http"
		req.URL.Host = origin.Host
		req.Host = origin.Host
	}
	proxy := &httputil.ReverseProxy{Director: director}
	if config.Host.Live {
		fmt.Printf("Routing to host: %v\n", origin.Host)
	} else {
		fmt.Printf("Serving mock data from: %v\n", origin.Host)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if config.Host.Live {
			_directory = r.URL.Path
			proxy.ServeHTTP(NewCustomWriter(w), r)
		} else {
			serveMockData(NewCustomWriter(w), r)
		}
	})

	log.Fatal(http.ListenAndServe(":9000", nil))
}
