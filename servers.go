package main

import (
	"fmt"
	"net/http"
)

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

func InitServers(c *Config) {
	for _, server := range c.Servers {
		s := Server{
			port: server.Port,
			s:    http.NewServeMux(),
		}
		go http.ListenAndServe(fmt.Sprintf(":%d", server.Port), s)
	}
}
