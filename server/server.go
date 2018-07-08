package server

import (
	"database/sql"
	"net/http"

	"github.com/MangoHacks/Mango2019-API/database"
	"github.com/gorilla/mux"
)

// server represents an instance of the server and all the dependencies required across it.
type server struct {
	db     *sql.DB
	router *mux.Router
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

func (s *server) bindHandlers() {
	s.router.HandleFunc("/preregister", handlePreregister())
}

// StartServer starts a new server.
func StartServer() error {
	s, err := newServer()
	if err != nil {
		return err
	}
	s.bindHandlers()
	if err := http.ListenAndServe(":9000", s.router); err != nil {
		return err
	}
	return nil
}
