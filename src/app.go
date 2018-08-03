package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"path/filepath"
)

type TestGenieConfig struct {
	Host struct {
		Url  string `yaml:"url"`
		Live string `yaml:"live"`
	}
}

func get_config() TestGenieConfig {
	config := TestGenieConfig{}

	filename, _ := filepath.Abs("./config.yaml")
	yamlFile, err1 := ioutil.ReadFile(filename)

	if err1 != nil {
		log.Fatalf("error: %v", err1)
	}

	err2 := yaml.Unmarshal(yamlFile, &config)

	if err2 != nil {
		log.Fatalf("error: %v", err2)
	}

	return config
}

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
		fmt.Printf("%v\n", w)
		proxy.ServeHTTP(w, r)
	})

	log.Fatal(http.ListenAndServe(":9000", nil))
}
