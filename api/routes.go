package api

import "github.com/go-chi/chi"

func (s *Server) registrationRoutes() {
	s.baseHttp.Mux.Route("/v1", func(v1 chi.Router) {
		v1.Get("/specification", s.endpoints.Specification())
		v1.Post("/mint_token", s.endpoints.Mint())
	})
}
