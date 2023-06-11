package xzcert

import (
	xzcommon "github.com/averyyan/xz/core/common"
	xzcontext "github.com/averyyan/xz/core/context"
)

func GetService(name xzcommon.ServiceName) *service {
	if v := xzcontext.GetService(CertServiceType, name); v != nil {
		if res, ok := v.(*service); ok {
			return res
		}
	}
	return nil
}
