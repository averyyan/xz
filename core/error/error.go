package xzerror

// 错误接口
type ErrorInterface interface {
	Error() string

	Code() int

	Message() string

	Details() []interface{}
}

type ErrorBuilder interface {
	New(code int, msg string, details ...interface{}) ErrorInterface

	NewCustom() ErrorInterface
}
