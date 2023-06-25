package main

import (
	"github.com/igdotog/core/config"
	"github.com/igdotog/core/http"
	"github.com/igdotog/core/logger"
	"github.com/igefined/mint/api"
	"github.com/igefined/mint/api/endpoints"
	"github.com/igefined/mint/use"

	"github.com/go-playground/validator/v10"
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
			use.NewComposite,
		),

		use.Constructor(),

		fx.Invoke(
			func(_ *api.Server) {},
		),
	)

	application.Run()
}
