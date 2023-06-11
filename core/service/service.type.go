package xzservice

import (
	"context"

	xzcommon "github.com/averyyan/xz/core/common"
)

type RegFunc func(raw []byte) map[xzcommon.ServiceName]Service

// Service 接口
type Service interface {
	// Service 的启动函数
	Bootstrap(context.Context)

	// Service 的结束函数
	Interrupt(context.Context)

	// 获取Service的名称
	GetName() xzcommon.ServiceName

	// 获取Service的类型
	GetType() xzcommon.ServiceType

	// 获取Service的详情
	GetDescription() string

	// 用于打印Service信息
	String() string
}
