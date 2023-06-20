package main

import (
	"log"
	"net/http"
	"sync"
)

func main() {
	go InitServers()

	lb := Loadbalancer{
		servers:       []string{"http://localhost:8001", "http://localhost:8002"},
		currentServer: 0,
		mux:           new(sync.RWMutex),
	}

	http.HandleFunc("/", lb.ServeHTTP)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
