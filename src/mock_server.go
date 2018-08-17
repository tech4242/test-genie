package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"path"
)

func serveMockData(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("%s", r.URL)
	filePath := path.Join("./responses", r.URL.Path, "response.json")
	file, e := ioutil.ReadFile(filePath)
	if e != nil {
		fmt.Printf("Non-cached request")
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "500 - Something bad happened!")
	} else {
		customWriter := NewCustomWriter(w)
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(customWriter, "%s", string(file))
	}
}
