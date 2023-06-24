package endpoints

import "net/http"

// Mint godoc
// @summary Mint token
// @tags Mint
// @success 200 {object} http.Empty
// @router /v1/mint_token [post]
func (e endpoints) Mint() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(http.StatusOK)
	}
}
