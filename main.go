package main

import (
	"io"
	"log"
	"net/http"
)

type Loadbalancer struct {
	servers       []string
	currentServer int
}

func (lb *Loadbalancer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	res, _ := http.Get(lb.servers[lb.currentServer])
	lb.currentServer = (lb.currentServer + 1) % len(lb.servers)
	io.Copy(w, res.Body)
}

type Server struct {
	port int
	s    *http.ServeMux
}

func (s Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Helo from " + r.Host + "\n"))
}

func main() {
	s1 := Server{
		port: 8001,
		s:    http.NewServeMux(),
	}

	s2 := Server{
		port: 8002,
		s:    http.NewServeMux(),
	}

	go func() {
		s1.s.HandleFunc("/", s1.ServeHTTP)
		log.Fatal(http.ListenAndServe(":8001", s1))
	}()

	go func() {
		s2.s.HandleFunc("/", s2.ServeHTTP)
		log.Fatal(http.ListenAndServe(":8002", s2))
	}()

	lb := Loadbalancer{
		servers:       []string{"http://localhost:8001", "http://localhost:8002"},
		currentServer: 0,
	}

	http.HandleFunc("/", lb.ServeHTTP)

	log.Fatal(http.ListenAndServe(":8000", nil))
}
