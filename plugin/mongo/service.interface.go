package xzmongo

import (
	"context"
	"crypto/tls"
	"fmt"

	xzcommon "github.com/averyyan/xz/core/common"
	xzlogger "github.com/averyyan/xz/core/logger"
	xzregister "github.com/averyyan/xz/core/register"
	"go.mongodb.org/mongo-driver/mongo"
)

// Service 的启动函数
func (s *service) Bootstrap(_ context.Context) {
	s.bootstrapOnce.Do(func() {
		if s.certEntry != nil {
			s.opts.TLSConfig = &tls.Config{
				Certificates:       make([]tls.Certificate, 0),
				InsecureSkipVerify: s.insecureSkipVerify,
			}
			if s.certEntry.GetCertificate() != nil {
				s.opts.TLSConfig.Certificates = append(s.opts.TLSConfig.Certificates, *s.certEntry.GetCertificate())
			}

			if s.certEntry.GetRootCA() != nil {
				s.opts.TLSConfig.Certificates = append(s.opts.TLSConfig.Certificates, *s.certEntry.GetRootCACertificate())
			}
		}
		// connect to mongo
		if client, err := mongo.Connect(context.Background(), s.opts); err != nil {
			xzregister.ShutdownWithError(err)
		} else {
			xzlogger.Info(s, fmt.Sprintf("Creating mongoDB client at %v success", s.opts.Hosts))
			s.client = client
		}
		// try ping
		pingCtx, cancel := context.WithTimeout(context.Background(), s.pingTimeout)
		defer cancel()
		if err := s.client.Ping(pingCtx, nil); err != nil {
			xzlogger.Error(s, fmt.Sprintf("Ping mongoDB at %v failed", s.opts.Hosts))
			xzregister.ShutdownWithError(err)
		}
		// create database
		for k, v := range s.mongoDbOpts {
			s.mongoDbMap[k] = s.client.Database(k, v...)
			xzlogger.Info(s, fmt.Sprintf("Creating database instance [%s] success", k))
		}
	})
}

// Service 的结束函数
func (s *service) Interrupt(_ context.Context) {
	s.client.Disconnect(context.Background())
}

// 获取Service的名称
func (s *service) GetName() xzcommon.ServiceName {
	return s.serviceName
}

// 获取Service的类型
func (s *service) GetType() xzcommon.ServiceType {
	return MongoServiceType
}

// 获取Service的详情
func (s *service) GetDescription() string {
	return s.serviceDescription
}

// 用于打印Service信息
func (s *service) String() string {
	return ""
}
