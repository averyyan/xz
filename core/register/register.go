package xzregister

import (
	"fmt"

	xzcommon "github.com/averyyan/xz/core/common"
	xzlogger "github.com/averyyan/xz/core/logger"
	xzservice "github.com/averyyan/xz/core/service"
)

func New(opts ...Option) *Register {
	register := &Register{}
	for _, opt := range opts {
		opt(register)
	}
	// 默认注册一个stdOut logger
	serviceMap := make(map[xzcommon.ServiceName]xzservice.Service)
	loggerService := xzlogger.NewLoggerStdoutService()
	serviceMap[loggerService.GetName()] = loggerService
	register.registerService(serviceMap)
	return register
}

func RegisterService(register *Register, regFunc xzservice.RegFunc) {
	register.registerServiceByFunc(regFunc)
}

func ShutdownWithError(err error) {
	if err == nil {
		err = fmt.Errorf("关闭服务:%s", err)
	}
	panic(err)
}
