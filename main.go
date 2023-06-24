package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/igdotog/core/config"
	"github.com/igdotog/core/http"
	"github.com/igdotog/core/logger"
	"github.com/igefined/mint/api"
	"github.com/igefined/mint/api/endpoints"

	"go.uber.org/fx"
)

// @title           Mint internal API
// @version         1.0

// @contact.name   API Documentation
// @contact.url    ig.pomazkov@gmail.com

// @accept json
// @produce json
// @schemes http
func main() {
	ctx := config.SigTermIntCtx()
	log := logger.New()

	cfg, err := config.New(ctx, &config.Http{})
	if err != nil {
		log.Fatal(err)
	}

	application := fx.New(
		fx.Supply(cfg, log, validator.New()),

		fx.Provide(
			api.NewServer,
			http.NewBaseServer,
			endpoints.New,
		),

		fx.Invoke(
			func(_ *api.Server) {},
		),
	)

	application.Run()
}
