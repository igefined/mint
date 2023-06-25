package endpoints

import (
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"

	"github.com/igdotog/core/http/respond"
)

type (
	mintReq struct {
		Receiver string `json:"receiver"`
		Amount   string `json:"amount" validate:"numeric"`
	}

	mintResp struct {
		TxHash string `json:"tx_hash"`
	}
)

// Mint godoc
// @summary Mint token
// @tags Mint
// @success 200 {object} mintResp
// @param account body mintReq true "mint structure"
// @param Private-Key header string true "Private key for signing tx"
// @router /v1/mint_token [post]
func (e endpoints) Mint() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req *mintReq

		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			respond.InternalError(w, err)

			return
		}

		if err := e.validator.StructCtx(r.Context(), req); err != nil {
			respond.InternalError(w, err)

			return
		}

		amount := new(big.Int)
		if _, err := fmt.Sscan(req.Amount, amount); err != nil {
			respond.InternalError(w, err)

			return
		}

		txHash, err := e.use.Minter().Mint(r.Context(), req.Receiver, amount)
		if err != nil {
			respond.InternalError(w, err)

			return
		}

		resp := &mintResp{
			TxHash: txHash,
		}

		respond.Successfully(w, resp)
	}
}
