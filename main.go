package main

import (
	"log"
	"net/http"
	"sync"
)

func main() {
	c := ReadConfig()
	InitServers(c)

	lb := Loadbalancer{
		servers:       c.Servers,
		currentServer: 0,
		mux:           new(sync.RWMutex),
		protocol:      c.Protocol,
	}

	http.HandleFunc("/", lb.ServeHTTP)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
