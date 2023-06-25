package minter

import (
	"context"
	"fmt"
	"math/big"

	"github.com/igefined/mint/contracts"
	"github.com/igefined/mint/utils/uctx"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

func (s service) Mint(ctx context.Context, address string, amount *big.Int) (hash string, err error) {
	token, err := contracts.NewToken(common.HexToAddress(s.cfg.Address), s.client)
	if err != nil {
		err = fmt.Errorf("minter.Mint: error create a new instance of Token, bound to a specific deployed contract: %v", err)

		return
	}

	chainId, err := s.client.ChainID(ctx)
	if err != nil {
		err = fmt.Errorf("minter.Mint: error get chain id: %v", err)

		return
	}

	gasPrice, err := s.client.SuggestGasPrice(ctx)
	if err != nil {
		err = fmt.Errorf("minter.Mint: error get suggest gas price: %v", err)

		return
	}

	opt, err := bind.NewKeyedTransactorWithChainID(uctx.GetPrivateKey(ctx), chainId)
	if err != nil {
		err = fmt.Errorf("minter.Mint: error create tx: %v", err)

		return
	}
	opt.GasPrice = gasPrice

	mint, err := token.Mint(opt, common.HexToAddress(address), amount)
	if err != nil {
		err = fmt.Errorf("minter.Mint: error mint tokens: %v", err)

		return
	}

	hash = mint.Hash().Hex()

	return
}
