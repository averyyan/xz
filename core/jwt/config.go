package xzjwt

import (
	xzcommon "github.com/averyyan/xz/core/common"
	xzservice "github.com/averyyan/xz/core/service"
	xzutil "github.com/averyyan/xz/core/util"
)

func RegisterServiceYAML(raw []byte) map[xzcommon.ServiceName]xzservice.Service {
	var config Config
	serviceMap := make(map[xzcommon.ServiceName]xzservice.Service)
	if err := xzutil.UnmarshalConfigYAML(raw, &config); err != nil {
		panic(err.Error())
	}
	for i := range config.Jwt {
		element := config.Jwt[i]
		if !element.Enabled {
			continue
		}
		service := RegisterService(
			WithName(xzcommon.ServiceName(element.Name)),
			WithDescription(element.Description),
			WithSymmetric(element.Symmetric),
			WithIgnore(element.Ignore),
		)
		serviceMap[service.serviceName] = service
	}
	return serviceMap
}

// 注册Entry
func RegisterService(opts ...Option) *service {
	entry := &service{
		serviceName:        "Jwt服务",
		serviceDescription: "Jwt服务",
	}
	for i := range opts {
		opts[i](entry)
	}
	return entry
}
