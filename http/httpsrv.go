package http

import (
	"byex.io/irisrv/basic"
	"byex.io/irisrv/http/controller"
	"byex.io/irisrv/public"
	"context"
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/recover"
	"strconv"
	"time"
)

const (
	DefaultHttpPort = 8900
)

type SrvOptions struct {
	Address string
	Port    string
}

func (option *SrvOptions) ListenAddr() (addr string) {
	if len(option.Port) <= 0 {
		option.Port = strconv.Itoa(DefaultHttpPort)
	}

	if len(option.Address) <= 0 {
		addr = ":" + option.Port
	}

	return
}

type IrisSrv struct {
	options SrvOptions
	jwt     basic.JwtOptions
	app     *iris.Application
}

func NewHttpSrv(jwtOption basic.JwtOptions, option SrvOptions) *IrisSrv {
	httpSrv := &IrisSrv{
		options: option,
		jwt:     jwtOption,
	}

	app := iris.New()
	httpSrv.app = app
	app.Use(recover.New())
	app.RegisterView(iris.HTML("./views", ".html"))

	iris.RegisterOnInterrupt(func() {
		timeout := 5 * time.Second
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()
		// close all hosts
		app.Shutdown(ctx)
	})

	public.InitYaag(app)
	basic.JWTAuth(app, httpSrv.jwt)
	basic.InitBasicCtrl(app)
	Initialize()

	return httpSrv
}

func (s *IrisSrv) Start() {
	go s.app.Run(iris.Addr(s.options.ListenAddr()), iris.WithoutInterruptHandler)
}

func Initialize() {
	basic.RegisterCtrl(&controller.CCoinController{})
	basic.RegisterCtrl(&controller.FCoinController{})
	basic.RegisterCtrl(&controller.DemoController{})
}
