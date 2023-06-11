package xzregister

import (
	xzutil "github.com/averyyan/xz/core/util"
)

type Option func(reg *Register)

// 设置配置文件路径
func WithConfigPath(filePath string) Option {
	return func(reg *Register) {
		raw, err := xzutil.ReadYAML(filePath)
		if err != nil {
			panic(err)
		}
		reg.configRaw = raw
	}
}

// 设置配置文件数据流
func WithConfigRaw(raw []byte) Option {
	return func(reg *Register) {
		if raw != nil {
			reg.configRaw = raw
			return
		}
		panic("raw 未存在")
	}
}
