package httpserver

import (
	"fmt"
	"net/http"
)

type HttpServer struct{}

func NewHttpServer() *HttpServer {
	return &HttpServer{}
}

func (s *HttpServer) Listen(addr string, handler http.Handler) error {
	fmt.Println("http server is running on", addr)
	return http.ListenAndServe(addr, handler)
}
