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
	Debug   bool
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
	options    SrvOptions
	jwtOptions basic.JwtOptions
	jwt        iris.Handler
	app        *iris.Application
}

func NewHttpSrv(jwtOption basic.JwtOptions, option SrvOptions) *IrisSrv {
	app := iris.New()
	app.Use(recover.New())
	app.RegisterView(iris.HTML("./views", ".html"))

	iris.RegisterOnInterrupt(func() {
		timeout := 5 * time.Second
		ctx, cancel := context.WithTimeout(context.Background(), timeout)
		defer cancel()
		// close all hosts
		app.Shutdown(ctx)
	})

	httpSrv := &IrisSrv{
		options:    option,
		jwtOptions: jwtOption,
		jwt:        basic.JWTAuth(jwtOption),
		app:        app,
	}
	return httpSrv
}

func (s *IrisSrv) Start() {
	go s.app.Run(iris.Addr(s.options.ListenAddr()), iris.WithoutInterruptHandler)
}

func (s *IrisSrv) BeforeStart() {
	if s.options.Debug {
		public.InitYaag(s.app)
	}

	basic.InitBasicCtrl(s.app, s.jwt, "/v1")

	basic.RegisterCtrl(new(controller.CCoinController))
	basic.RegisterCtrl(new(controller.FCoinController))
	basic.RegisterCtrl(new(controller.DemoController))
}
