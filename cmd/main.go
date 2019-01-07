package main

import (
	"byex.io/httpsrv/node"
	"fmt"
	"github.com/bxsmart/bxcore/log"
	"gopkg.in/urfave/cli.v1"
	"os"
	"os/signal"
	"path/filepath"
	"reflect"
	"runtime"
)

func main() {
	app := newApp()
	app.Action = startNode
	app.HideVersion = true
	app.Copyright = "Copyright 2017-2022 The Byex Authors"
	globalFlags := globalFlags()
	app.Flags = append(app.Flags, globalFlags...)

	app.Before = func(ctx *cli.Context) error {
		runtime.GOMAXPROCS(runtime.NumCPU())
		return nil
	}

	app.After = func(ctx *cli.Context) error {
		fmt.Println(ctx.Args())
		return nil
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

}

func startNode(ctx *cli.Context) error {
	globalConfig := setGlobalConfig(ctx)

	logger := log.Initialize(globalConfig.Log)
	defer func() {
		if nil != logger {
			logger.Sync()
		}
	}()

	var n *node.Node
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	signal.Notify(signalChan, os.Kill)
	go func() {
		for {
			select {
			case sig := <-signalChan:
				log.Infof("captured %s, exiting...\n", sig.String())
				if nil != n {
					n.BeforeStop()

					n.Stop()

					n.AfterStart()
				}
				os.Exit(1)
			}
		}
	}()

	n = node.NewNode(logger, globalConfig)

	n.BeforeStart()

	n.Start()

	n.AfterStart()

	log.Info(">>>>    order srv started !!!    <<<<")

	n.Wait()
	return nil
}

func newApp() *cli.App {
	app := cli.NewApp()
	app.Name = filepath.Base(os.Args[0])
	app.Version = Version
	app.Usage = "BYEX Command Line interface"
	app.Author = ""
	app.Email = ""
	return app
}

func globalFlags() []cli.Flag {
	return []cli.Flag{
		cli.StringFlag{
			Name:  "config,c",
			Usage: "config file",
		},
	}
}

func setGlobalConfig(ctx *cli.Context) *node.GlobalConfig {
	file := ""
	if ctx.IsSet("config") {
		file = ctx.String("config")
	}
	globalConfig := node.LoadConfig(file)

	if _, err := node.Validator(reflect.ValueOf(globalConfig).Elem()); nil != err {
		panic(err)
	}

	return globalConfig
}
