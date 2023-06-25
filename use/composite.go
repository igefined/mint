package use

import (
	"github.com/igefined/mint/use/minter"

	"go.uber.org/fx"
)

func Constructor() fx.Option {
	return fx.Provide(
		minter.New,
	)
}

type UseCase interface {
	Minter() minter.Service
}

type use struct {
	minter minter.Service
}

func NewComposite(minter minter.Service) UseCase {
	return &use{minter: minter}
}

func (u use) Minter() minter.Service {
	return u.minter
}
