package main

import (
	"opsmicro/src/book/auth/handler"
	"book/auth/model"
	"book/basic"
	"book/basic/config"
	"fmt"
	"github.com/micro/cli"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/consul"
	"github.com/micro/go-micro/util/log"
	"time"

	auth "book/auth/proto/auth"
)

func main() {

	// 初始化配置、数据库等信息
	basic.Init()
	// 使用consul注册
    micReg := consul.NewRegistry(registryOptions)
	// New Service
	service := micro.NewService(
		micro.Name("demo.micro.book.srv.auth"),
		micro.Registry(micReg),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init(
		micro.Action(func(c *cli.Context) {
			// 初始化模型层
			model.Init()
			// 初始化handler
			handler.Init()
		}),
		)

	// Register Handler
	auth.RegisterServiceHandler(service.Server(), new(handler.Service))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

func registryOptions(ops *registry.Options) {
	consulCfg := config.GetConsulConfig()
	ops.Timeout = time.Second * 5
	ops.Addrs = []string{fmt.Sprintf("%s:%d", consulCfg.GetHost(), consulCfg.GetPort())}
}