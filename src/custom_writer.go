package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

var _directory = ""

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
		log.Fatalf("error: %v", err)
	}
	return c.ResponseWriter.Write(data)
}

func (c *customWriter) WriteHeader(i int) {
	c.ResponseWriter.WriteHeader(i)
}
