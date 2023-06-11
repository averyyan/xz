package xzjwt

import (
	xzcommon "github.com/averyyan/xz/core/common"
	xzservice "github.com/averyyan/xz/core/service"
	xzutil "github.com/averyyan/xz/core/util"
	"github.com/golang-jwt/jwt/v5"
)

const JwtServiceType xzcommon.ServiceType = "JwtService"

type Service interface {
	xzservice.Service
	VerifyIgnore(path string) bool
	ParseToken(tokenString string) (*jwt.Token, error)
	GeneratedRegisteredClaimsToken(claims jwt.Claims) (string, error)
}

type service struct {
	serviceName        xzcommon.ServiceName
	serviceDescription string
	symmetric          SymmetricConfig
	ignore             []string
}

func (s *service) VerifyIgnore(path string) bool {
	return xzutil.InArrayString(path, s.ignore)
}

func (s *service) ParseToken(tokenString string) (*jwt.Token, error) {
	return s.symmetric.ParseToken(tokenString)
}

func (s *service) GeneratedRegisteredClaimsToken(claims jwt.Claims) (string, error) {
	return s.symmetric.GeneratedRegisteredClaimsToken(claims)
}
