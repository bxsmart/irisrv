package basic

import (
	"github.com/kataras/iris"
	irisctx "github.com/kataras/iris/context"
	"reflect"
	"strings"
)

type Controller interface {
	Init(jwt iris.Handler) Controller
	Version(ver string) Controller
	Router(app *iris.Application)
}

type basicController struct {
	App *iris.Application
	Jwt iris.Handler
	Ver string

	CtrlKv map[string]interface{}
}

var basicCtrl basicController

func InitBasicCtrl(app *iris.Application, jwt iris.Handler, ver string) {
	app.Use(customerLogger())
	app.Use(Before)
	app.Done(After)

	app.OnErrorCode(iris.StatusNotFound, notFound)
	app.OnErrorCode(iris.StatusInternalServerError, internalServerError)

	basicCtrl = basicController{
		App:    app,
		Jwt:    jwt,
		Ver:    ver,
		CtrlKv: make(map[string]interface{}),
	}
}

func notFound(ctx iris.Context) {
	ctype := ctx.GetHeader(irisctx.ContentTypeHeaderKey)
	if ctx.IsAjax() || strings.Contains(ctype, irisctx.ContentJSONHeaderValue) {
		ctx.WriteString("404 Not Found !!!")
	} else {
		ctx.View("errors/404.html")
	}
}

func internalServerError(ctx iris.Context) {
	ctype := ctx.GetHeader(irisctx.ContentTypeHeaderKey)
	if ctx.IsAjax() || strings.Contains(ctype, irisctx.ContentJSONHeaderValue) {
		ctx.WriteString("Ops something went wrong, try again")
	} else {
		ctx.View("errors/500.html")
	}
}

type BasicController struct {
	Controller
	Ver string
	Jwt iris.Handler
}

func RegisterCtrl(c Controller) Controller {
	rtype := reflect.TypeOf(c)
	println("#### register: ", rtype.String())

	c.Version(basicCtrl.Ver)
	c.Init(basicCtrl.Jwt)
	c.Router(basicCtrl.App)

	basicCtrl.CtrlKv[rtype.String()] = c

	return c
}

func (c *BasicController) Version(ver string) Controller {
	c.Ver = ver
	return c
}

func (c *BasicController) Init(jwt iris.Handler) Controller {
	c.Jwt = jwt
	return c
}

func (c *BasicController) GetValue(ctx iris.Context, key string) string {
	value := ctx.Params().Get(key)
	if len(value) <= 0 {
		value = ctx.FormValue(key)
	}

	if len(value) <= 0 {
		value = ctx.PostValue(key)
	}

	return value
}
