package main

import (
	"context"
	"io"
	"time"

	"omo-msa-startkit/config"

	proto "omo-msa-startkit/proto/startkit"

	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/client"
	"github.com/micro/go-micro/v2/logger"
	_ "github.com/micro/go-plugins/registry/consul/v2"
	_ "github.com/micro/go-plugins/registry/etcdv3/v2"
)

func main() {
	config.Setup()
	service := micro.NewService(
		micro.Name("omo.msa.startkit.tester"),
	)
	service.Init()

	cli := service.Client()
	cli.Init(
		client.Retries(3),
		client.RequestTimeout(time.Second*1),
		client.Retry(func(_ctx context.Context, _req client.Request, _retryCount int, _err error) (bool, error) {
			if nil != _err {
				logger.Errorf("%v | [ERR] retry %d, reason is %v\n\r", time.Now().String(), _retryCount, _err)
				return true, nil
			}
			return false, nil
		}),
	)

	echo := proto.NewEchoService("omo.msa.startkit", cli)

	logger.Trace("----------------------------------------------------------")
	// Call
	{
		rsp, err := echo.Call(context.Background(), &proto.Request{
			Name: time.Now().String() + " | MSA-StartKit",
		})
		if err != nil {
			logger.Error(err)
		} else {
			logger.Info(rsp.Msg)
		}
	}

	stream, err := echo.PingPong(context.Background())
	if err != nil {
		logger.Error(err)
		return
	}
	defer stream.Close()
	stroke := int64(0)
	for range time.Tick(1 * time.Second) {
		stroke = stroke + 1
		stream.Send(&proto.Ping{Stroke: stroke})
		rsp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			logger.Error(err)
			continue
		}
		logger.Infof("Pong %v", rsp.Stroke)
	}
}
