package main

import (
	"context"
	"fmt"
	"time"

	"omo-msa-startkit/config"
	msa "omo-msa-startkit/proto/msa"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/client"
	_ "github.com/micro/go-plugins/registry/consul/v2"
	_ "github.com/micro/go-plugins/registry/etcdv3/v2"
)

func main() {
	config.Setup()
	service := micro.NewService()
	service.Init()

	cli := service.Client()
	cli.Init(
		client.Retries(3),
		client.RequestTimeout(time.Second*1),
		client.Retry(func(_ctx context.Context, _req client.Request, _retryCount int, _err error) (bool, error) {
			if nil != _err {
				fmt.Println(fmt.Sprintf("%v | [ERR] retry %d, reason is %v\n\r", time.Now().String(), _retryCount, _err))
				return true, nil
			}
			return false, nil
		}),
	)

	startkit := msa.NewStartKitService("omo.msa.startkit", cli)

	for range time.Tick(4 * time.Second) {
		fmt.Println("----------------------------------------------------------")
		// Make request
		rsp, err := startkit.Call(context.Background(), &msa.Request{
			Name: time.Now().String() + " | MSA-StartKit",
		})
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(rsp.Msg)
		}
	}
}
