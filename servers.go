package main

import (
	"fmt"
	"net/http"
)

type Server struct {
	port int
	s    *http.ServeMux
}

func (s Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from " + r.Host + " " + r.URL.Path + " " + r.Method + "\n"))
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
