package main

import (
	"log"
	"net/http"
)

func InitServers() {
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
}
