package xzcontext

import (
	"os"
	"time"

	xzcommon "github.com/averyyan/xz/core/common"
	xzservice "github.com/averyyan/xz/core/service"
)

// 单例模式实现 保证全局变量安全
func GetGlobalAppCtx() *appContext {
	once.Do(func() {
		appCtx = &appContext{
			startTime:   time.Now().Local(),
			services:    make([]xzservice.Service, 0),
			servicesMap: make(map[xzcommon.ServiceType]map[xzcommon.ServiceName]int),
			shutdownSig: make(chan os.Signal),
		}
	})
	return appCtx
}

// 获取所有的服务
func GetServices() []xzservice.Service {
	return GetGlobalAppCtx().services
}

// 注册服务到全局上下文 返回服务index
func RegisterService(serviceType xzcommon.ServiceType, serviceName xzcommon.ServiceName, service xzservice.Service) {
	GetGlobalAppCtx().services = append(GetGlobalAppCtx().services, service)
	if GetGlobalAppCtx().servicesMap[serviceType] == nil {
		GetGlobalAppCtx().servicesMap[serviceType] = make(map[xzcommon.ServiceName]int)
	}
	GetGlobalAppCtx().servicesMap[serviceType][serviceName] = len(GetGlobalAppCtx().services) - 1
}

func GetServicesByType(serviceType xzcommon.ServiceType) map[xzcommon.ServiceName]xzservice.Service {
	m := make(map[xzcommon.ServiceName]xzservice.Service)
	for _, index := range GetGlobalAppCtx().servicesMap[serviceType] {
		serv := GetGlobalAppCtx().services[index]
		m[serv.GetName()] = serv
	}
	return m
}

// 获取特定服务
func GetService(serviceType xzcommon.ServiceType, serviceName xzcommon.ServiceName) xzservice.Service {
	if GetGlobalAppCtx().servicesMap[serviceType] == nil {
		return nil
	}
	return GetGlobalAppCtx().services[GetGlobalAppCtx().servicesMap[serviceType][serviceName]]
}
