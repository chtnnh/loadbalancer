package main

import (
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
)

func TestRace(t *testing.T) {
	c := ReadConfig()
	InitServers(c)

	lb := Loadbalancer{
		servers:       c.Servers,
		currentServer: 0,
		mux:           new(sync.RWMutex),
		protocol:      c.Protocol,
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
