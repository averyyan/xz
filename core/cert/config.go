package xzcert

import (
	xzcommon "github.com/averyyan/xz/core/common"
	xzservice "github.com/averyyan/xz/core/service"
	xzutil "github.com/averyyan/xz/core/util"
)

func RegisterServiceYAML(raw []byte) map[xzcommon.ServiceName]xzservice.Service {
	var config Config
	entryMap := make(map[xzcommon.ServiceName]xzservice.Service)
	if err := xzutil.UnmarshalConfigYAML(raw, &config); err != nil {
		panic(err.Error())
	}
	for i := range config.Cert {
		element := config.Cert[i]
		if !element.Enabled {
			continue
		}
		entry := registerService(
			WithName(xzcommon.ServiceName(element.Name)),
			WithDescription(element.Description),
			WithCaPath(element.CAPath),
			WithCertPemPath(element.CertPemPath),
			WithKeyPemPath(element.KeyPemPath),
		)
		entryMap[entry.serviceName] = entry
	}
	return entryMap
}

func registerService(opts ...Option) *service {
	entry := &service{
		serviceName:        "cert服务",
		serviceDescription: "cert服务",
	}
	for i := range opts {
		opts[i](entry)
	}
	return entry
}
