package main

import (
	"fmt"
	"net/http"
)

var _directory = ""

type customWriter struct {
	http.ResponseWriter
	status      int
	wroteHeader bool
}

func NewCustomWriter(w http.ResponseWriter) *customWriter {
	return &customWriter{ResponseWriter: w}
}
func (w *customWriter) Status() int {
	return w.status
}

func (c *customWriter) Header() http.Header {
	return c.ResponseWriter.Header()
}

func (c *customWriter) Write(data []byte) (int, error) {
	if c.status == http.StatusOK {
		parentFolder := "responses"
		fileName := "response"
		fileExtension := ".json"
		writeToFileSystem(data, _directory, fileName, fileExtension, parentFolder)

	}
	fmt.Printf("Data: %v\n", string(data))
	return c.ResponseWriter.Write(data)
}

func (c *customWriter) WriteHeader(i int) {

	if c.wroteHeader == true {
		return
	}
	c.status = i
	c.wroteHeader = true
	c.ResponseWriter.WriteHeader(i)

}
