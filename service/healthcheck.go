package service

import (
	"log"
	"net"
	"net/url"
	"time"
)

func ping(url *url.URL) bool {
	conn, err := net.DialTimeout("tcp", url.Host, time.Second*3)
	if err != nil {
		return false
	}
	defer conn.Close()
	return true
}

func PerformHealthCheck() {
	t := time.NewTicker(time.Second * 10)
	for range t.C {
		for _, s := range GlobalLoadBalancer.Servers {
			url, err := url.Parse(s.Addr)
			if err != nil {
				log.Fatal(err.Error())
			}
			isAlive := ping(url)
			if !isAlive {
				log.Printf("Unreachable to %v", url.Host)
				s.mu.Lock()
				s.IsAlive = false
				s.mu.Unlock()
			}
		}
	}
}
