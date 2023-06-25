package uctx

import (
	"context"
	"crypto/ecdsa"
)

type ContextKey int

const (
	keyPrivateKey ContextKey = iota
)

func PutPrivateKey(ctx context.Context, pvk *ecdsa.PrivateKey) context.Context {
	return context.WithValue(ctx, keyPrivateKey, pvk)
}

func GetPrivateKey(ctx context.Context) (pvk *ecdsa.PrivateKey) {
	if value := ctx.Value(keyPrivateKey); value == nil {
		return
	} else {
		pvk, _ = value.(*ecdsa.PrivateKey)
	}

	return
}
