package xzcommon

type ctxKey int

const (
	EventID ctxKey = iota
	JwtToken
)
