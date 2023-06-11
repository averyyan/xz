package test

import (
	"context"
	"os"
	"testing"

	xzcert "github.com/averyyan/xz/core/cert"
	xzjwt "github.com/averyyan/xz/core/jwt"
	xzregister "github.com/averyyan/xz/core/register"
	xzmongo "github.com/averyyan/xz/plugin/mongo"
)

func TestMain(m *testing.M) {
	ctx := context.Background()
	reg := xzregister.New(xzregister.WithConfigPath("boot.yaml"))

	xzregister.RegisterService(reg, xzcert.RegisterServiceYAML)
	xzregister.RegisterService(reg, xzjwt.RegisterServiceYAML)
	xzregister.RegisterService(reg, xzmongo.RegisterServiceYAML)

	reg.Bootstrap(ctx)
	os.Exit(m.Run())
}
