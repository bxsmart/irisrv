package controller

import (
	"byex.io/httpsrv/basic"
	"github.com/kataras/iris"
)

type CCoinController struct {
	basic.BasicController
}

func (c *CCoinController) Router(app *iris.Application) {
	ccoin := app.Party(c.Ver)

	ccoin.Any("/ccoin/info", c.CCoinInfo)
}

// @pa
func (c *CCoinController) CCoinInfo(ctx iris.Context) {
	id := c.GetValue(ctx, "id")

	ctx.Writef("########" + string(id))
	ctx.Next()
}
