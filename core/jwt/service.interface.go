package xzjwt

import (
	"context"

	xzcommon "github.com/averyyan/xz/core/common"
)

// Service 的启动函数
func (s *service) Bootstrap(_ context.Context) {
}

// Service 的结束函数
func (s *service) Interrupt(_ context.Context) {
}

// 获取Service的名称
func (s *service) GetName() xzcommon.ServiceName {
	return s.serviceName
}

// 获取Service的类型
func (s *service) GetType() xzcommon.ServiceType {
	return JwtServiceType
}

// 获取Service的详情
func (s *service) GetDescription() string {
	return s.serviceDescription
}

// 用于打印Service信息
func (s *service) String() string {
	return ""
}
