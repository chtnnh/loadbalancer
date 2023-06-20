package main

import (
	"log"
	"net/http"
)

func main() {
	go InitServers()

	lb := Loadbalancer{
		servers:       []string{"http://localhost:8001", "http://localhost:8002"},
		currentServer: 0,
	}

	http.HandleFunc("/", lb.ServeHTTP)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
