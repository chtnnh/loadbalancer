package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Config struct {
	LB       URI         `json:"lb"`
	Servers  []ServerURI `json:"servers"`
	Protocol Algo        `json:"protocol"`
}

type Algo int

const (
	RoundRobin Algo = iota
	WeightedRoundRobin
)

type URI struct {
	Host string `json:"host"`
	Port int    `json:"port"`
}

type ServerURI struct {
	Host   string `json:"host"`
	Port   int    `json:"port"`
	Weight int    `json:"weight"`
	Uri    string `json:"uri"`
}

func ReadConfig() *Config {
	f, err := os.Open("config.json")
	if err != nil {
		panic(fmt.Errorf("%w", err))
	}

	defer f.Close()

	b, _ := ioutil.ReadAll(f)
	c := &Config{}

	err = json.Unmarshal(b, c)
	if err != nil {
		panic(fmt.Errorf("%w", err))
	}

	servers := []ServerURI{}
	for _, server := range c.Servers {
		if server.Weight < 1 {
			panic(fmt.Errorf("server weight cannot be less than 1"))
		}
		for i := 0; i < server.Weight; i += 1 {
			servers = append(servers, server)
		}
	}
	c.Servers = servers

	return c
}
