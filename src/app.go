package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func main() {
	config := get_config()
	origin, _ := url.Parse(config.Host.Url)
	director := func(req *http.Request) {
		req.URL.Scheme = "http"
		req.URL.Host = origin.Host
		req.Host = origin.Host
	}

	fmt.Printf("Routing to host: %v\n", origin.Host)

	proxy := &httputil.ReverseProxy{Director: director}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_directory = r.URL.Path
		proxy.ServeHTTP(NewCustomWriter(w), r)
	})

	log.Fatal(http.ListenAndServe(":9000", nil))
}
