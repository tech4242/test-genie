package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strings"
)

var _directory = ""

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

type customWriter struct {
	http.ResponseWriter
}

func NewCustomWriter(w http.ResponseWriter) *customWriter {
	return &customWriter{w}
}

func (c *customWriter) Header() http.Header {
	return c.ResponseWriter.Header()
}

func (c *customWriter) Write(data []byte) (int, error) {
	fmt.Println(string(data))       //get response here
	fmt.Println(string(_directory)) //get response here

	err := ioutil.WriteFile("src/responses/"+strings.Replace(_directory, "/", "-", -1)+".json", data, 0777)
	if err != nil {
		panic(err)
	}
	return c.ResponseWriter.Write(data)
}

func (c *customWriter) WriteHeader(i int) {
	c.ResponseWriter.WriteHeader(i)
}
