package xzlogger

import (
	xzcommon "github.com/averyyan/xz/core/common"
	xzservice "github.com/averyyan/xz/core/service"
	"go.uber.org/zap"
)

const LoggerServiceType xzcommon.ServiceType = "LoggerService"

type Service interface {
	xzservice.Service
	IsDefault() bool
}

type service struct {
	*zap.Logger
	serviceName        xzcommon.ServiceName
	serviceDescription string
	isDefault          bool
}

func (s *service) IsDefault() bool {
	return s.isDefault
}
