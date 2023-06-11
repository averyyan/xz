package xzmongo

import (
	"time"

	xzcert "github.com/averyyan/xz/core/cert"
	xzcommon "github.com/averyyan/xz/core/common"
	mongoOpt "go.mongodb.org/mongo-driver/mongo/options"
)

type Option func(s *service)

func WithName(name xzcommon.ServiceName) Option {
	return func(s *service) {
		s.serviceName = name
	}
}

func WithDescription(description string) Option {
	return func(s *service) {
		s.serviceDescription = description
	}
}

func WithCertEntry(name string) Option {
	return func(s *service) {
		if len(name) > 0 {
			certEntry := xzcert.GetService(xzcommon.ServiceName(name))
			if certEntry == nil {
				panic("certEntry未存在")
			} else {
				s.certEntry = certEntry
			}
		}
	}
}

func WithClientOptions(opt *mongoOpt.ClientOptions) Option {
	return func(s *service) {
		if opt != nil {
			s.opts = opt
		}
	}
}

func WithPingTimeout(tout int) Option {
	return func(s *service) {
		if tout > 0 {
			s.pingTimeout = time.Duration(tout) * time.Millisecond
		}
	}
}

func WithInsecureSkipVerify(skip bool) Option {
	return func(s *service) {
		s.insecureSkipVerify = skip
	}
}

func WithDatabase(dbName string, dbOpts ...*mongoOpt.DatabaseOptions) Option {
	return func(s *service) {
		if _, ok := s.mongoDbOpts[dbName]; !ok {
			s.mongoDbOpts[dbName] = make([]*mongoOpt.DatabaseOptions, 0)
		}
		s.mongoDbOpts[dbName] = append(s.mongoDbOpts[dbName], dbOpts...)
	}
}
