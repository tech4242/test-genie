package main

import "log"

func errorHandler(err error) {
	if err != nil {
		log.Fatalf("error: %v", err)
	}
}
