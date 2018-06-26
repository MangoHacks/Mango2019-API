package server

import (
	"github.com/MangoHacks/Mango2019-API/handlers"
)

func (s *Server) bindRoutes() {
	s.Router.HandleFunc("/preregister", handlers.HandlePreregister())
}
