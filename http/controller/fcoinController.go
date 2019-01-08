package controller

import (
	"byex.io/irisrv/basic"
	"github.com/kataras/iris"
)

type FCoinController struct {
	basic.BasicController
}

func (c *FCoinController) Router(app *iris.Application) {
	fcoin := app.Party(c.Ver)

	fcoin.Any("/fcoin/position", c.FCoinPosition)
}

func (c *FCoinController) FCoinPosition(ctx iris.Context) {
	id := c.GetValue(ctx, "id")

	ctx.Writef("########" + string(id))
	ctx.Next()
}
