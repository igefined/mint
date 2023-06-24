package api

import (
	"github.com/igdotog/core/http"
	"github.com/igdotog/core/logger"
	"github.com/igefined/mint/api/endpoints"
)

type Server struct {
	log       *logger.Logger
	baseHttp  *http.BaseServer
	endpoints endpoints.Endpoints
}

func NewServer(
	log *logger.Logger,
	baseHttp *http.BaseServer,
	endpoints endpoints.Endpoints,
) *Server {
	instance := &Server{
		log:       log,
		baseHttp:  baseHttp,
		endpoints: endpoints,
	}

	instance.registrationRoutes()

	return instance
}
