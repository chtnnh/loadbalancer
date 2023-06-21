package main

import (
	"fmt"
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

	var res *http.Response

	switch method := r.Method; method {
	case http.MethodGet:
		res, _ = http.Get(fmt.Sprintf("%s%s", lb.servers[lb.currentServer].Uri, r.URL.Path))
	case http.MethodPost:
		res, _ = http.Post(
			fmt.Sprintf("%s%s", lb.servers[lb.currentServer].Uri, r.URL.Path),
			"",
			r.Body,
		)
	default:
		panic(fmt.Errorf("method %s not supported", method))
	}

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
