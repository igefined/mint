package middleware

import (
	"errors"
	"net/http"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/igdotog/core/http/respond"
	"github.com/igefined/mint/utils/uctx"
)

var ErrMissingPrivateKey = errors.New("missing private key request header Private-Key")

func PrivateKey() func(handler http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			privateKeyAsString, err := PrivateKeyExtractor(r)
			if err != nil {
				respond.Forbidden(w, err)

				return
			}

			privateKey, err := crypto.HexToECDSA(privateKeyAsString)
			if err != nil {
				respond.Forbidden(w, ErrMissingPrivateKey)

				return
			}

			ctx := uctx.PutPrivateKey(r.Context(), privateKey)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func PrivateKeyExtractor(r *http.Request) (string, error) {
	const authKey = "Private-Key"

	keyFromHeader := r.Header.Get(authKey)
	if len(keyFromHeader) == 0 {
		return "", ErrMissingPrivateKey
	}

	return keyFromHeader, nil
}
