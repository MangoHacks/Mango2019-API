// Package server deals with the initialization of the server as
// well as the listening and handling of resources.
package server

import (
	"net/http"

	"github.com/MangoHacks/Mango2019-API/database"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

// Server represents an instance of the server and all the dependencies required across it.
type Server struct {
	router *mux.Router
	db     *database.DB
}

// New constructs a new server.
//
// The server includes a mux.Router pointer,
// and an sql.DB pointer
func New() (*Server, error) {
	r := mux.NewRouter()
	db, err := database.New("postgres")
	if err != nil {
		return nil, err
	}
	return &Server{
		router: r,
		db:     db,
	}, nil
}

// bindHandlers sets the handler functions of the router.
// This allows the web service to listen to and handle requests made to these
// resources.
func (s *Server) bindHandlers() {
	s.router.HandleFunc("/preregistration", handlePreregistration(s.db))
	s.router.HandleFunc("/registration", handleRegistration(s.db))
}

// Start starts a server.
func (s *Server) Start() error {
	s.bindHandlers()
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"POST", "OPTIONS", "GET"},
	})
	return http.ListenAndServe(":9000", c.Handler(s.router))
}
