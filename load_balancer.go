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
	res, _ := http.Get(lb.servers[lb.currentServer].Uri)
	defer lb.next()
	io.Copy(w, res.Body)
}
