package web

import (
	"fmt"
	"log"
	"net/http"
)

// Server represents the web server.
type Server struct {
	Addr   string
	Router *http.Server
}

// NewServer creates a new web server instance.
func NewServer(addr string, router *http.Server) *Server {
	return &Server{
		Addr:   addr,
		Router: router,
	}
}

// Start starts the web server.
func (s *Server) Start() {
	log.Printf("Server listening on %s", s.Addr)
	if err := s.Router.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Failed to start server: %v", err)
	}
}

// Stop gracefully stops the web server.
func (s *Server) Stop() {
	if err := s.Router.Shutdown(nil); err != nil {
		log.Fatalf("Failed to stop server: %v", err)
	}
	fmt.Println("Server stopped")
}
