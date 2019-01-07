package controller

import (
	"byex.io/httpsrv/basic"
	"github.com/kataras/iris"
)

type DemoController struct {
	basic.BasicController
}

type myXML struct {
	Result string `xml:"result"`
}

func (c *DemoController) Router(app *iris.Application) {
	demo := app.Party(c.Ver)

	demo.Any("/demo/json", c.Json)
	demo.Any("/demo/text", c.Text)
	demo.Any("/demo/xml", c.Xml)
	demo.Any("/demo/complex", c.Complex)
}

func (c *DemoController) Json(ctx iris.Context) {
	ctx.JSON(iris.Map{"result": "Hello World!"})
	ctx.Next()
}

func (c *DemoController) Text(ctx iris.Context) {
	ctx.Text("Hello World!")
	ctx.Next()
}

func (c *DemoController) Xml(ctx iris.Context) {
	ctx.XML(myXML{Result: "Hello World!"})
	ctx.Next()
}

func (c *DemoController) Complex(ctx iris.Context) {
	value := c.GetValue(ctx, "key")

	ctx.JSON(iris.Map{"value": value})
	ctx.Next()
}
