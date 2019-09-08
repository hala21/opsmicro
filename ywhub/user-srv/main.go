package main

import (
	"fmt"
	"github.com/micro/cli"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/consul"
	"github.com/micro/go-micro/util/log"
	"opsmicro/ywhub/basic"
	"opsmicro/ywhub/basic/config"
	"opsmicro/ywhub/user-srv/handler"
	"opsmicro/ywhub/user-srv/model"
	user "opsmicro/ywhub/user-srv/proto/user"
	"time"
	//"opsmicro/ywhub/user-srv/subscriber"
)

func main() {
	// 初始化配置、数据库等信息
	basic.Init()
	// 使用consul注册
	micReg := consul.NewRegistry(registryOptions)
	// New Service
	service := micro.NewService(
		micro.Name("ywhub.micro.v1.basic.srv.user"),
		micro.Registry(micReg),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init(micro.Action(func(context *cli.Context) {
		model.Init()
		handler.Init()
	}))

	// Register Handler
	_ = user.RegisterUserHandler(service.Server(), new(handler.User))

	// Register Struct as Subscriber
	//micro.RegisterSubscriber("ywhub.micro.basic.srv.user.srv.user", service.Server(), new(subscriber.User))

	// Register Function as Subscriber
	//micro.RegisterSubscriber("ywhub.micro.basic.srv.user.srv.user", service.Server(), subscriber.Handler)

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
