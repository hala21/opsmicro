package main

import (
	"fmt"
	"github.com/micro/cli"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/consul"
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-micro/web"
	"net/http"
	"opsmicro/ywhub/basic"
	"opsmicro/ywhub/basic/config"
	"opsmicro/ywhub/user-web/handler"
	"time"
)

func main() {
	// basic init
	basic.Init()
	micReg := consul.NewRegistry(registryOptions)
	// create new web service
	service := web.NewService(
		web.Name("ywhub.micro.v1.basic.web.user"),
		web.Registry(micReg),
		web.Address(":8088"),
		web.Version("latest"),
	)

	// initialise service
	if err := service.Init(web.Action(
		func(c *cli.Context) {
			// 初始化handler
			handler.Init()
		})); err != nil {
		log.Fatal(err)
	}

	// 注册登录接口
	service.HandleFunc("/user/login", handler.Login)

	service.HandleFunc("/user/logout", handler.Logout)
	// register html handler
	service.Handle("/", http.FileServer(http.Dir("html")))

	// run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}

func registryOptions(ops *registry.Options) {
	consulCfg := config.GetConsulConfig()
	ops.Timeout = time.Second * 5
	ops.Addrs = []string{fmt.Sprintf("%s:%d", consulCfg.GetHost(), consulCfg.GetPort())}
}
