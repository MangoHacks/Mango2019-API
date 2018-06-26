package server

import (
	"github.com/gorilla/mux"
)

// Server represents an instance of the server and all the dependencies required across it.
type Server struct {
	DB     *string
	Router *mux.Router
	email  string
}

// New returns a new pointer to a server.
func New() *Server {
	var s Server
	r := mux.NewRouter()
	s.Router = r
	s.bindRoutes()
	return &s
}
