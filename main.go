package main

import (
	"github.com/micro/go-micro/util/log"
	"github.com/micro/go-micro"
	"github.com/alactic/shippydemo/handler"
	"github.com/alactic/shippydemo/subscriber"

	shippydemo "github.com/alactic/shippydemo/proto/shippydemo"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.shippydemo"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	shippydemo.RegisterShippydemoHandler(service.Server(), new(handler.Shippydemo))

	// Register Struct as Subscriber
	micro.RegisterSubscriber("go.micro.srv.shippydemo", service.Server(), new(subscriber.Shippydemo))

	// Register Function as Subscriber
	micro.RegisterSubscriber("go.micro.srv.shippydemo", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
