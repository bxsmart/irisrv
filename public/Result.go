package public

import (
	"github.com/kataras/iris"
	"net/http"
)

func Result(ctx iris.Context, code int, data interface{}, msg string) {
	ctx.StatusCode(http.StatusOK)
	ctx.JSON(iris.Map{"code": code, "data": data, "msg": msg})
}

func ResultOk(ctx iris.Context, data interface{}) {
	ctx.StatusCode(http.StatusOK)
	ctx.JSON(iris.Map{"code": http.StatusOK, "data": data, "msg": ""})
}

func ResultList(ctx iris.Context, data interface{}, total int64) {
	ctx.StatusCode(http.StatusOK)
	ctx.JSON(iris.Map{"code": http.StatusOK, "rows": data, "msg": "", "total": total})
}

func ResultOkMsg(ctx iris.Context, data interface{}, msg string) {
	ctx.StatusCode(http.StatusOK)
	ctx.JSON(iris.Map{"code": http.StatusOK, "data": data, "msg": msg})
}

func ResultFail(ctx iris.Context, err interface{}) {
	ctx.StatusCode(http.StatusOK)
	ctx.JSON(iris.Map{"code": http.StatusBadRequest, "data": nil, "msg": err})
}

func ResultFailData(ctx iris.Context, data interface{}, err interface{}) {
	ctx.StatusCode(http.StatusOK)
	ctx.JSON(iris.Map{"code": http.StatusBadRequest, "data": data, "msg": err})
}
