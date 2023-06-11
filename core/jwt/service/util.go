package xzjwtservice

import (
	"context"

	xzcommon "github.com/averyyan/xz/core/common"
	"github.com/golang-jwt/jwt/v5"
)

func GetTokenFromCtx(ctx context.Context) *jwt.Token {
	if token, ok := ctx.Value(xzcommon.JwtToken).(*jwt.Token); ok {
		return token
	}
	return nil
}

func GetRegisteredClaimsFromCtx(ctx context.Context) *jwt.RegisteredClaims {
	token := GetTokenFromCtx(ctx)
	if token == nil {
		return nil
	}
	if claims, ok := token.Claims.(*jwt.RegisteredClaims); ok {
		return claims
	}
	return nil
}
