package xzcert

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"

	xzcommon "github.com/averyyan/xz/core/common"
	xzregister "github.com/averyyan/xz/core/register"
	xzutil "github.com/averyyan/xz/core/util"
)

// Service 的启动函数
func (s *service) Bootstrap(_ context.Context) {
	s.bootstrapOnce.Do(func() {
		if len(s.keyPemPath) > 0 && len(s.certPemPath) > 0 {
			certPem, err := xzutil.ReadFile(s.certPemPath, s.embedFS)
			if err != nil {
				xzregister.ShutdownWithError(err)
			}
			keyPem, err := xzutil.ReadFile(s.keyPemPath, s.embedFS)
			if err != nil {
				xzregister.ShutdownWithError(err)
			}
			cert, err := tls.X509KeyPair(certPem, keyPem)
			if err != nil {
				xzregister.ShutdownWithError(err)
			}

			s.certificate = &cert
		}

		if len(s.caPath) > 0 {
			ca, err := xzutil.ReadFile(s.caPath, s.embedFS)
			if err != nil {
				xzregister.ShutdownWithError(err)
			}
			block, _ := pem.Decode(ca)
			if block == nil || block.Type != "CERTIFICATE" || len(block.Headers) != 0 {
				return
			}
			cert, err := x509.ParseCertificate(block.Bytes)
			if err != nil {
				xzregister.ShutdownWithError(err)
			}
			s.rootCA = cert
		}
	})

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
	return CertServiceType
}

// 获取Service的详情
func (s *service) GetDescription() string {
	return s.serviceDescription
}

// 用于打印Service信息
func (s *service) String() string {
	return ""
}
