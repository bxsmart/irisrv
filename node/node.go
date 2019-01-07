package node

import (
	"byex.io/httpsrv/basic"
	"byex.io/httpsrv/dao"
	"byex.io/httpsrv/mvc"
	"byex.io/httpsrv/public"
	"context"
	"fmt"
	"github.com/bxsmart/bxcore/cache"
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/recover"
	"go.uber.org/zap"
	"sync"
	"time"
)

type Node struct {
	globalConfig *GlobalConfig
	rdsSrv       *dao.RdsService
	wg           *sync.WaitGroup
	logger       *zap.Logger
	irisSrv      *iris.Application
}

func NewNode(logger *zap.Logger, globalConfig *GlobalConfig) *Node {
	n := &Node{}
	n.logger = logger
	n.globalConfig = globalConfig
	n.wg = new(sync.WaitGroup)

	// register
	n.registerMysql()
	n.registerCache()

	return n
}

func (n *Node) BeforeStart() {
}

func (n *Node) Start() {

	// 启动Rpc监听端口
	fmt.Println("step in ordsrv node start")
}

func (n *Node) AfterStart() {
	n.wg.Add(1)
}

func (n *Node) Wait() {
	n.wg.Wait()
}

func (n *Node) BeforeStop() {
}

// todo release zklock and kafka producers and consumers
func (n *Node) Stop() {
	n.wg.Done()
}

func (n *Node) AfterStop() {
}

func (n *Node) registerMysql() {
	n.rdsSrv = dao.NewDb(&n.globalConfig.Mysql)
}

func (n *Node) registerCache() {
	cache.NewCache(n.globalConfig.Redis)
}

func (n *Node) registerIrisSrv() {
	app := iris.New()
	n.irisSrv = app
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
	basic.JWTAuth(app, n.globalConfig.Jwt)
	basic.InitBasicCtrl(app)
	mvc.Initialize()

	app.Run(iris.Addr(":8000"), iris.WithoutInterruptHandler)
}
