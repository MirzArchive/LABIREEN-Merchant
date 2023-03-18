package jwtx

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type MerchantClaims struct {
	ID uuid.UUID `json:"id"`
	jwt.RegisteredClaims
}

func NewMerchantClaims(id uuid.UUID, exp time.Duration) MerchantClaims {
	return MerchantClaims{
		ID: id,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(exp)),
		},
	}
}
