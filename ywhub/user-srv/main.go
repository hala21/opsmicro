package main

import (
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/util/log"
	_ "opsmicro/ywhub/user-srv/handler"
	"opsmicro/ywhub/user-srv/subscriber"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("ywhub.micro.v1.basic.srv.user"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	user.RegisterUserHandler(service.Server(), new(opsmicro.User))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("ywhub.micro.basic.srv.user.srv.user", service.Server(), new(subscriber.User))

	// Register Function as Subscriber
	micro.RegisterSubscriber("ywhub.micro.basic.srv.user.srv.user", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
