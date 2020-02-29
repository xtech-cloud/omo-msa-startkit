package main

import (
	"context"
	"fmt"
	"time"

	"omo-msa-startkit/config"
	msa "omo-msa-startkit/proto/msa"

	"github.com/micro/go-micro"
	"github.com/micro/go-micro/v2/client"
	_ "github.com/micro/go-plugins/registry/consul/v2"
	_ "github.com/micro/go-plugins/registry/etcdv3/v2"
)

func main() {
	config.Setup()
	service := micro.NewService()

	/*
		// parse command line flags
		service.Init()

		// Use the generated client stub
		cl := msa.NewStartKitService("omo.msa.startkit", service.Client())
	*/

	cli := client.NewClient(
		client.Retry(func(_ctx context.Context, _req client.Request, _retryCount int, _err error) (bool, error) {
			if nil != _err {
				fmt.Println("[ERR] retry %d, reason is %v", _retryCount, _err)
				return true, nil
			}
			return false, nil
		}),
	)

	startkit := msa.NewStartKitService("omo.msa.startkit", cli)

	for range time.Tick(3 * time.Second) {
		fmt.Println("----------------------------------------------------------")
		// Make request
		rsp, err := startkit.Call(context.Background(), &msa.Request{
			Name: "John " + time.Now().String(),
		})
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(rsp.Msg)
	}
}
