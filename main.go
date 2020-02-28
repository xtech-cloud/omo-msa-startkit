package main

import (
	"omo-msa-startkit/config"
	"omo-msa-startkit/handler"

	_ "github.com/micro/go-plugins/registry/consul/v2"
	_ "github.com/micro/go-plugins/registry/kubernetes/v2"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/util/log"

	msa "omo-msa-startkit/proto/msa"
)

func main() {
	config.Setup()

	// New Service
	service := micro.NewService(
		micro.Name("omo.msa.startkit"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	msa.RegisterStartKitHandler(service.Server(), new(handler.StartKit))

	// Register Struct as Subscriber
	//micro.RegisterSubscriber("omo.msa.startkit", service.Server(), new(subscriber.StartKit))

	// Register Function as Subscriber
	//micro.RegisterSubscriber("omo.msa.startkit", service.Server(), subscriber.Handler)

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
