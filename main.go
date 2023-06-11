package main

import (
	"context"

	xzcert "github.com/averyyan/xz/core/cert"
	xzjwt "github.com/averyyan/xz/core/jwt"
	xzregister "github.com/averyyan/xz/core/register"
	xzmongo "github.com/averyyan/xz/plugin/mongo"
)

func main() {
	ctx := context.Background()
	reg := xzregister.New(xzregister.WithConfigPath("test/boot.yaml"))

	xzregister.RegisterService(reg, xzcert.RegisterServiceYAML)
	xzregister.RegisterService(reg, xzjwt.RegisterServiceYAML)
	xzregister.RegisterService(reg, xzmongo.RegisterServiceYAML)

	reg.Bootstrap(ctx)
	reg.WaitForShutdownSig(ctx)
}
