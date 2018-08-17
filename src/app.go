package main

import (
	"log"
	"net/http"
)

func main() {
	config := getConfig()

	if config.Host.Refresh {
		startCacheRefresh(config)
	} else {
		startReverseProxy(config)
	}

	log.Fatal(http.ListenAndServe(":9000", nil))
}
