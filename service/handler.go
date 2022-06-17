package service

import (
	"net/http"
	"net/http/httputil"

	"github.com/gin-gonic/gin"
)

// HandleReverseProxy transfers request to related server(s)
func HandleReverseProxy(c *gin.Context) {
	proxy := &httputil.ReverseProxy{
		Director: func(r *http.Request) {
			r.URL.Scheme = "http"
			r.URL.Host = GlobalLoadBalancer.SelectAvailableServer().Addr
		},
	}
	proxy.ServeHTTP(c.Writer, c.Request)
}
