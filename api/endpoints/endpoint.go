package endpoints

import (
	"net/http"

	"github.com/igdotog/core/logger"
	"github.com/igefined/mint/use"

	"github.com/go-playground/validator/v10"
)

type Endpoints interface {
	Mint() http.HandlerFunc
	Specification() http.HandlerFunc
}

type endpoints struct {
	log       *logger.Logger
	validator *validator.Validate

	use use.UseCase
}

func New(
	log *logger.Logger,
	validator *validator.Validate,
	use use.UseCase,
) Endpoints {
	return &endpoints{
		log:       log,
		validator: validator,
		use:       use,
	}
}
