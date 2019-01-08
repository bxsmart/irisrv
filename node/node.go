package node

import (
	"byex.io/irisrv/dao"
	"byex.io/irisrv/http"
	"github.com/bxsmart/bxcore/cache"
	"go.uber.org/zap"
	"sync"
)

type Node struct {
	globalConfig *GlobalConfig
	rdsSrv       *dao.RdsService
	wg           *sync.WaitGroup
	logger       *zap.Logger
	irisSrv      *http.IrisSrv
}

func NewNode(logger *zap.Logger, globalConfig *GlobalConfig) *Node {
	n := &Node{}
	n.logger = logger
	n.globalConfig = globalConfig
	n.wg = new(sync.WaitGroup)

	// register
	n.registerMysql()
	n.registerCache()
	n.registerIrisSrv()

	return n
}

func (n *Node) BeforeStart() {
	n.irisSrv.BeforeStart()
}

func (n *Node) Start() {
	n.irisSrv.Start()

	println("step in irisrv node start...")
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

func (n *Node) registerCache() {
	cache.NewCache(n.globalConfig.Redis)
}

func (n *Node) registerMysql() {
	n.rdsSrv = dao.NewDb(&n.globalConfig.Mysql)
}

func (n *Node) registerIrisSrv() {
	n.irisSrv = http.NewHttpSrv(n.globalConfig.Jwt, n.globalConfig.HttpSrv)
}
