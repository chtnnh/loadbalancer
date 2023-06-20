package main

import (
	"io"
	"net/http"
	"sync"
)

type Loadbalancer struct {
	servers       []string
	currentServer int
	mux           *sync.RWMutex
}

func (lb *Loadbalancer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	lb.mux.Lock()
	defer lb.mux.Unlock()
	res, _ := http.Get(lb.servers[lb.currentServer])
	defer lb.next()
	io.Copy(w, res.Body)
}

func (lb *Loadbalancer) next() {
	lb.currentServer = (lb.currentServer + 1) % len(lb.servers)
}

type Server struct {
	port int
	s    *http.ServeMux
}

func (s Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Helo from " + r.Host + "\n"))
}
