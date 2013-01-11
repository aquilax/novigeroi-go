package main

import (
	"log"
	"net/http"
)

type Server struct {
	port string
}

func NewServer() *Server {
	return &Server{":8080"}
}

func (s *Server) StartServer() {
	log.Printf("Starting server on port: %s\n", s.port)
	log.Fatal(http.ListenAndServe(s.port, nil))
}
