package xzjwt

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

func WithSymmetric(symmetric SymmetricConfig) Option {
	return func(s *service) {
		s.symmetric = symmetric
	}
}

func WithIgnore(ignore []string) Option {
	return func(s *service) {
		s.ignore = ignore
	}
}
