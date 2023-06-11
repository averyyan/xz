package xzcert

import xzcommon "github.com/averyyan/xz/core/common"

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

func WithCaPath(path string) Option {
	return func(s *service) {
		s.caPath = path
	}
}

func WithCertPemPath(path string) Option {
	return func(s *service) {
		s.certPemPath = path
	}
}

func WithKeyPemPath(path string) Option {
	return func(s *service) {
		s.keyPemPath = path
	}
}
