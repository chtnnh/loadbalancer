package main

import (
	"io"
	"net/http"
	"sync"
)

type Loadbalancer struct {
	servers       []ServerURI
	currentServer int
	mux           *sync.RWMutex
	protocol      Algo
}

func (lb *Loadbalancer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	lb.mux.Lock()
	defer lb.mux.Unlock()
	// TODO: replace http.Get with support for all methods
	res, _ := http.Get(lb.servers[lb.currentServer].Uri)
	defer lb.next()
	io.Copy(w, res.Body)
}

func (lb *Loadbalancer) next() {
	if lb.protocol == 0 {
		lb.roundRobin()
	}
}

func (lb *Loadbalancer) roundRobin() {
	lb.currentServer = (lb.currentServer + 1) % len(lb.servers)
}
