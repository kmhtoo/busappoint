package account

import (
	"github.com/kataras/iris"
	"github.com/kmhtoo/busappoint/app"
)

type AccountForget struct {
	*app.BusApp
}

func (c AccountForget) Serve(ctx *iris.Context) {
	ctx.Write("/account/forget")
}
