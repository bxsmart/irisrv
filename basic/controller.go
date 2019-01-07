package basic

import (
	"github.com/kataras/iris"
	irisctx "github.com/kataras/iris/context"
	"reflect"
	"strings"
)

type Controller interface {
	Router(app *iris.Application)
	Init()
	Version(ver string)
}

type basicController struct {
	App *iris.Application

	CtrlKv map[string]interface{}
}

var basicCtrl basicController

func InitBasicCtrl(app *iris.Application) {
	app.Use(customerLogger())
	app.Use(Before)
	app.Done(After)

	app.OnErrorCode(iris.StatusNotFound, notFound)
	app.OnErrorCode(iris.StatusInternalServerError, internalServerError)

	basicCtrl = basicController{App: app, CtrlKv: make(map[string]interface{})}
}

func RegisterCtrl(ctrl Controller) {
	rtype := reflect.TypeOf(ctrl)

	println(rtype.Name())

	ctrl.Version("/v1")
	ctrl.Router(basicCtrl.App)

	println("#### register ctrl: ", rtype.String())
	basicCtrl.CtrlKv[rtype.String()] = ctrl
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
}

func (c *BasicController) Version(ver string) {
	c.Ver = ver
}

func (c *BasicController) Init() {
	RegisterCtrl(c)
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
