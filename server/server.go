package server

import (
	"database/sql"
	"net/http"

	"github.com/MangoHacks/Mango2019-API/database"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

// server represents an instance of the server and all the dependencies required across it.
type server struct {
	router *mux.Router
	db     *sql.DB
}

// newServer constructs a new server.
func newServer() (*server, error) {
	r := mux.NewRouter()
	db, err := database.New()
	if err != nil {
		return nil, err
	}
	return &server{
		router: r,
		db:     db,
	}, nil
}

// bindHandlers sets the handler functions of the router.
func (s *server) bindHandlers() {
	s.router.HandleFunc("/preregister", handlePreregister(s.db))
	s.router.HandleFunc("/register", handleRegister(s.db))
}

// StartServer starts a new server.
func StartServer() error {
	s, err := newServer()
	if err != nil {
		return err
	}
	s.bindHandlers()
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "OPTIONS", "DELETE"},
	})
	if err := http.ListenAndServe(":9000", c.Handler(s.router)); err != nil {
		return err
	}
	return nil
}
