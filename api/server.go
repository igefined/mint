package api

import (
	"net/http"

	server "github.com/igdotog/core/http"
	httpMiddleware "github.com/igdotog/core/http/middleware"
	"github.com/igdotog/core/logger"
	"github.com/igefined/mint/api/endpoints"

	"github.com/go-chi/chi/middleware"
)

type Server struct {
	log       *logger.Logger
	baseHttp  *server.BaseServer
	endpoints endpoints.Endpoints
}

func NewServer(
	log *logger.Logger,
	baseHttp *server.BaseServer,
	endpoints endpoints.Endpoints,
) *Server {
	instance := &Server{
		log:       log,
		baseHttp:  baseHttp,
		endpoints: endpoints,
	}

	instance.useInterceptors()
	instance.registrationRoutes()

	return instance
}

func (a *Server) headersSet(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-Frame-Options", "deny")
		w.Header().Set("X-XSS-Protection", "1; mode=block")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.Header().Set("Content-Type", "application/json")

		next.ServeHTTP(w, r)
	})
}

func (a *Server) useInterceptors() {
	a.baseHttp.Mux.Use(
		httpMiddleware.JsonLogger(a.log, []string{"/status"}),
		middleware.Recoverer,
		a.headersSet,
	)
}
