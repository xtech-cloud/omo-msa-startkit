package main

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"omo-msa-startkit/config"
	"omo-msa-startkit/handler"
	"os"
	"path/filepath"
	"time"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/logger"

	proto "omo-msa-startkit/proto/startkit"
)

func main() {
	config.Setup()

	// New Service
	service := micro.NewService(
		micro.Name("omo.msa.startkit"),
		micro.Version("latest"),
		micro.RegisterTTL(time.Second*time.Duration(config.Schema.Service.TTL)),
		micro.RegisterInterval(time.Second*time.Duration(config.Schema.Service.Interval)),
		micro.Address(config.Schema.Service.Address),
	)

	// Initialise service
	service.Init()

	// Register Handler
	proto.RegisterEchoHandler(service.Server(), new(handler.Echo))

	// Register Struct as Subscriber
	//micro.RegisterSubscriber("omo.msa.startkit", service.Server(), new(subscriber.StartKit))

	// Register Function as Subscriber
	//micro.RegisterSubscriber("omo.msa.startkit", service.Server(), subscriber.Handler)

	app, _ := filepath.Abs(os.Args[0])

	logger.Info("-------------------------------------------------------------")
	logger.Info("- Micro Service Agent -> Run")
	logger.Info("-------------------------------------------------------------")
	logger.Infof("- version      : %s", BuildVersion)
	logger.Infof("- application  : %s", app)
	logger.Infof("- md5          : %s", md5hex(app))
	logger.Infof("- build        : %s", BuildTime)
	logger.Infof("- commit       : %s", CommitID)
	logger.Info("-------------------------------------------------------------")
	// Run service
	if err := service.Run(); err != nil {
		logger.Fatal(err)
	}
}

func md5hex(_file string) string {
	h := md5.New()

	f, err := os.Open(_file)
	if err != nil {
		return ""
	}
	defer f.Close()

	io.Copy(h, f)

	return hex.EncodeToString(h.Sum(nil))
}
