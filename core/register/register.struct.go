package xzregister

import (
	"context"
	"fmt"
	"os"
	"runtime/debug"

	xzcommon "github.com/averyyan/xz/core/common"
	xzcontext "github.com/averyyan/xz/core/context"
	xzlogger "github.com/averyyan/xz/core/logger"
	xzservice "github.com/averyyan/xz/core/service"
)

// 注册器
type Register struct {
	configRaw []byte // 配置文件流
}

// 监听程序中断事件
func (register *Register) WaitForShutdownSig(ctx context.Context) {
	xzcontext.GetGlobalAppCtx().WaitForShutdownSig()
	register.interrupt(ctx)
}

// 注册服务
func (register *Register) registerService(serviceMap map[xzcommon.ServiceName]xzservice.Service) {
	defer syncLog()
	for _, service := range serviceMap {
		serviceType := service.GetType()
		serviceName := service.GetName()
		// 注册到appContext
		xzcontext.RegisterService(serviceType, serviceName, service)
	}
}

// 核心函数 显式注册Entry
func (register *Register) registerServiceByFunc(regFunc xzservice.RegFunc) {
	defer syncLog()
	if len(register.configRaw) == 0 {
		panic("配置文件未存在")
	}
	register.registerService(regFunc(register.configRaw))
}

// 核心函数 服务启动函数
func (register *Register) Bootstrap(ctx context.Context) {
	defer syncLog()
	// 顺序启动
	services := xzcontext.GetServices()
	for _, serv := range services {
		serv.Bootstrap(ctx)
		xzlogger.Info(serv, "服务已启动")
	}
}

// 核心函数 服务中断执行
func (register *Register) interrupt(ctx context.Context) {
	defer syncLog()
	services := xzcontext.GetServices()
	for i := range services {
		serv := services[len(services)-1-i]
		serv.Interrupt(ctx)
		xzlogger.Info(serv, "服务已结束")
	}
}

// 用于在启动、注册、中断函数中的Panic错误
func syncLog() {
	if r := recover(); r != nil {
		fmt.Printf("服务初始化错误, 关闭服务 %s \n%s", r.(string), string(debug.Stack()))
		os.Exit(1)
	}
}
