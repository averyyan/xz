package xzcert

import (
	"crypto/tls"
	"crypto/x509"
	"embed"
	"sync"

	xzcommon "github.com/averyyan/xz/core/common"
	xzservice "github.com/averyyan/xz/core/service"
)

type Service interface {
	xzservice.Service
	SetEmbedFS(fs *embed.FS)
	GetRootCA() *x509.Certificate
	GetRootCACertificate() *tls.Certificate
	GetCertificate() *tls.Certificate
}

const CertServiceType xzcommon.ServiceType = "CertService"

type service struct {
	serviceName        xzcommon.ServiceName
	serviceDescription string
	caPath             string
	keyPemPath         string
	certPemPath        string
	embedFS            *embed.FS
	rootCA             *x509.Certificate
	certificate        *tls.Certificate
	bootstrapOnce      sync.Once
}

func (s *service) SetEmbedFS(fs *embed.FS) {
	s.embedFS = fs
}

func (s *service) GetRootCA() *x509.Certificate {
	return s.rootCA
}

// 转化到Certificate
func (s *service) GetRootCACertificate() *tls.Certificate {
	return &tls.Certificate{Certificate: [][]byte{s.rootCA.Raw}}
}

func (s *service) GetCertificate() *tls.Certificate {
	return s.certificate
}
