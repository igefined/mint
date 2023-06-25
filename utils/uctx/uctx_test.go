//go:build units

package uctx

import (
	"context"
	"testing"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/assert"
)

func TestGetPrivateKey(t *testing.T) {
	t.Run("success, private key", func(t *testing.T) {
		privateKey, err := crypto.GenerateKey()
		assert.NoError(t, err)

		ctx := PutPrivateKey(context.Background(), privateKey)

		pvk := GetPrivateKey(ctx)
		assert.NoError(t, err)
		assert.Equal(t, privateKey, pvk)
	})
}
