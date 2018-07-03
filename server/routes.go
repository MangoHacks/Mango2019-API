package server

func (s *Server) bindRoutes() {
	s.router.HandleFunc("/preregister", handlePreregister())
}
