package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/GrokkingSystemDesign/LoadBalancer/service"
	"github.com/gin-gonic/gin"
)

func gracefulShutdown() {
	s := make(chan os.Signal, 1)
	signal.Notify(s, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-s
		log.Println("load balancer had been shutdown.")
		os.Exit(0)
	}()
}

func main() {

	go service.PerformHealthCheck()
	go gracefulShutdown()

	router := gin.Default()
	router.SetTrustedProxies(nil)
	router.POST("/gateway/v1", service.HandleReverseProxy)
	router.Run(service.GlobalLoadBalancer.Addr)
}
