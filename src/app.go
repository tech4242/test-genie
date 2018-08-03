package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {
	const (
		defaultRedirect      = "http://127.0.0.1:8000"
		defaultRedirectUsage = "default redirect url, 'http://127.0.0.1:8000'"
	)

	redirecturl := flag.String("redirect-url", defaultRedirect, defaultRedirectUsage)
	flag.Parse()

	origin, _ := url.Parse(*redirecturl)

	director := func(req *http.Request) {
		req.URL.Scheme = "http"
		req.URL.Host = origin.Host
		req.Host = origin.Host
	}

	fmt.Printf("Routing to host: %v\n", origin.Host)

	proxy := &httputil.ReverseProxy{Director: director}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		proxy.ServeHTTP(w, r)
	})

	log.Fatal(http.ListenAndServe(":9000", nil))
}
