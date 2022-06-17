package service

import (
	"sync"
)

var GlobalLoadBalancer = LoadBalancer{
	Addr: ":8080",
	Servers: []*Server{
		{Addr: "http://localhost:8081/", IsAlive: true},
		{Addr: "http://localhost:8082/", IsAlive: true},
		{Addr: "http://localhost:8083/", IsAlive: false},
		{Addr: "http://localhost:8084/", IsAlive: true},
	},
}

type Server struct {
	// Addr the address with which to access the server
	Addr string
	// IsAlive true if the server is alive and able to serve requests
	IsAlive bool
	mu sync.Mutex
}

type LoadBalancer struct {
	Addr    string
	Servers []*Server
	mu      sync.Mutex
	rrCount uint32 // Round Robin
}

func (lb *LoadBalancer) SelectAvailableServer() *Server {
	lb.mu.Lock()
	defer lb.mu.Unlock()

	s := lb.Servers[lb.rrCount%uint32(len(lb.Servers))]
	for !s.IsAlive {
		lb.rrCount++
		s = lb.Servers[lb.rrCount%uint32(len(lb.Servers))]
	}
	lb.rrCount++
	return s
}
