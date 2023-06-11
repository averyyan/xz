package xzjwt

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

type SymmetricConfig struct {
	Algorithm string `yaml:"algorithm" json:"algorithm"` // 加密算法
	Secret    string `yaml:"secret" json:"secret"`       // 加密私钥
}

// 生成 RegisteredClaims Token
func (conf *SymmetricConfig) GeneratedRegisteredClaimsToken(claims jwt.Claims) (string, error) {
	var alg jwt.SigningMethod
	switch conf.Algorithm {
	case "HS256":
		alg = jwt.SigningMethodHS256
	case "HS384":
		alg = jwt.SigningMethodHS384
	default:
		return "", fmt.Errorf("算法【%s】未支持", conf.Algorithm)
	}
	return jwt.NewWithClaims(alg, claims).SignedString([]byte(conf.Secret))
}

// 序列化Token
func (conf *SymmetricConfig) ParseToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		if t.Method.Alg() != conf.Algorithm {
			return nil, fmt.Errorf("jwt加密算法类型错误: %v", t.Header["alg"])
		}
		return []byte(conf.Secret), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}
