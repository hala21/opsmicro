package main

import (
	"book/user-srv/basic"
	"book/user-srv/basic/config"
	"book/user-srv/handler"
	"book/user-srv/model"
	"fmt"
	"github.com/micro/cli"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/consul"
	"github.com/micro/go-micro/util/log"
	"time"

	user "book/user-srv/proto/user"
)

func main() {
	// 初始化数据库和配置信息
	basic.Init()

	// 使用consul注册
	microReg := consul.NewRegistry(registryOptions)

	// New Service
	service := micro.NewService(
		micro.Name("demo.micro.book.srv.user"),
		micro.Registry(microReg),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init(
		micro.Action(func(c *cli.Context) {
			model.Init()
			handler.Init()
		}),
	)

	// Register Handler
	user.RegisterUserHandler(service.Server(), new(handler.Service))

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
