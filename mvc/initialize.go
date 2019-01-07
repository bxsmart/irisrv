package mvc

import (
	"byex.io/httpsrv/basic"
	"byex.io/httpsrv/mvc/controller"
	"byex.io/httpsrv/mvc/entity"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	"github.com/kataras/iris"
)

func Initialize() {
	basic.RegisterCtrl(&controller.CCoinController{})
	basic.RegisterCtrl(&controller.FCoinController{})
	basic.RegisterCtrl(&controller.DemoController{})
}

func InitRdsEngine(app *iris.Application) {
	engine, err := xorm.NewEngine("sqlite3", "./data/httpsrv.db")
	if err != nil {
		app.Logger().Fatalf("orm failed to initialized: %v", err)
	}

	iris.RegisterOnInterrupt(func() {
		engine.Close()
	})

	engine.SetColumnMapper(core.GonicMapper{})

	err = engine.Sync2(new(entity.FCoinAccount), new(entity.CCoinAccount))
}
