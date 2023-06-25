package api

import (
	"github.com/go-chi/chi"

	"github.com/igefined/mint/api/middleware"
)

func (s *Server) registrationRoutes() {
	s.baseHttp.Mux.Route("/v1", func(v1 chi.Router) {
		v1.Get("/specification", s.endpoints.Specification())
		v1.With(middleware.PrivateKey()).Post("/mint_token", s.endpoints.Mint())
	})
}
