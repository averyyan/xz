package xzjwt

import (
	"context"
	"fmt"
	"strings"

	xzcommon "github.com/averyyan/xz/core/common"
	xzgrpcerr "github.com/averyyan/xz/core/error/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// 将 JWT 配置转化为 grpc.UnaryServerInterceptor
func (s *service) ToGrpcUnaryInterceptor() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (resp interface{}, err error) {
		// 是否免验证接口
		if s.VerifyIgnore(info.FullMethod) {
			return handler(ctx, req)
		}
		// 获取Token
		md := getIncomingHeaders(ctx)
		authorization := getFirstHeader(md, xzcommon.HeaderAuthorization)
		authArr := strings.Split(authorization, " ")
		authType := authArr[0]
		tokenString := authArr[1]
		switch authType {
		case "Bearer":
			token, err := s.ParseToken(tokenString)
			if err != nil {
				return nil, xzgrpcerr.Unauthenticated(err.Error()).Err()
			}
			ctx = context.WithValue(ctx, xzcommon.JwtToken, token)
		default:
			return nil, xzgrpcerr.Unauthenticated(fmt.Sprintf("验证类型【%s】未存在", authType)).Err()
		}
		return handler(ctx, req)

	}
}

func getIncomingHeaders(ctx context.Context) metadata.MD {
	// called from server
	if v, ok := metadata.FromIncomingContext(ctx); ok {
		return v
	}

	return metadata.Pairs()
}

func getFirstHeader(md metadata.MD, key string) string {
	headers := md.Get(key)

	if len(headers) > 0 {
		return headers[0]
	}

	return ""
}
