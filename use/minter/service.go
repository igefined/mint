package minter

import (
	"context"
	"log"
	"math/big"
	"time"

	"github.com/igdotog/core/logger"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/heetch/confita"
	"github.com/heetch/confita/backend/env"
)

type Service interface {
	Mint(ctx context.Context, address string, amount *big.Int) (string, error)
}

type (
	service struct {
		logger *logger.Logger
		cfg    *config
		client *ethclient.Client
	}

	config struct {
		Url     string `config:"MINTER_CLIENT_URL,required"`
		Address string `config:"MITER_CONTRACT_ADDRESS,required"`
	}
)

const defaultClientUrl = "https://data-seed-prebsc-1-s1.binance.org:8545/"

func New(logger *logger.Logger) (Service, error) {
	cfg := config{
		Url: defaultClientUrl,
		// Address
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	if err := confita.NewLoader(env.NewBackend()).Load(ctx, &cfg); err != nil {
		log.Fatalf("minter: failed to load config: %s", err)
	}

	client, err := ethclient.Dial(cfg.Url)
	if err != nil {
		return nil, err
	}

	return &service{
		logger: logger,
		cfg:    &cfg,
		client: client,
	}, nil
}
