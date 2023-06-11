package xzcontext

import (
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	xzcommon "github.com/averyyan/xz/core/common"
	xzservice "github.com/averyyan/xz/core/service"
)

// 全局服务储存
var appCtx *appContext
var once sync.Once

// 主要用于服务储存
type appContext struct {
	startTime   time.Time                                             // 开始时间记录
	services    []xzservice.Service                                   // 全局服务储存
	servicesMap map[xzcommon.ServiceType]map[xzcommon.ServiceName]int // 服务地址查询
	shutdownSig chan os.Signal                                        // 系统信号
}

func init() {
	// 注册系统信号监控
	signal.Notify(
		GetGlobalAppCtx().shutdownSig,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
	)
}

// 监控服务结束信号
func (ctx *appContext) WaitForShutdownSig() {
	<-ctx.shutdownSig
}
