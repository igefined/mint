package endpoints

import (
	"net/http"

	"github.com/igdotog/core/logger"

	"github.com/go-playground/validator/v10"
)

type Endpoints interface {
	Mint() http.HandlerFunc
	Specification() http.HandlerFunc
}

type endpoints struct {
	log       *logger.Logger
	validator *validator.Validate
}

func New(
	log *logger.Logger,
	validator *validator.Validate,
) Endpoints {
	return &endpoints{
		log:       log,
		validator: validator,
	}
}
