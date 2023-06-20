package main

import (
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
)

func TestRace(t *testing.T) {
	InitServers()

	lb := Loadbalancer{
		servers:       []string{"http://localhost:8001", "http://localhost:8002"},
		currentServer: 0,
	}

	server := httptest.NewServer(http.HandlerFunc(lb.ServeHTTP))
	defer server.Close()

	wg := &sync.WaitGroup{}
	for i := 0; i < 50; i += 1 {
		wg.Add(1)
		go func() {
			defer wg.Done()

			req, err := http.NewRequest(http.MethodGet, server.URL, nil)
			if err != nil {
				panic(err)
			}

			res, err := http.DefaultClient.Do(req)
			if err != nil {
				panic(err)
			}

			defer res.Body.Close()
		}()
	}
	wg.Wait()
}
