package main

import (
	"context"
	"omo-msa-startkit/config"
	"testing"

	msa "omo-msa-startkit/proto/msa"

	micro "github.com/micro/go-micro/v2"
	_ "github.com/micro/go-plugins/registry/consul/v2"
)

var service msa.StartKitService

func TestInit(_t *testing.T) {
	config.Setup()
	svc := micro.NewService(micro.Name("omo.msa.startkit.client"))
	svc.Init()
	service = msa.NewStartKitService("omo.msa.startkit", svc.Client())
	_t.Log("init finish")
}

func TestCall(_t *testing.T) {
	rsp, err := service.Call(context.Background(), &msa.Request{Name: "MSA"})
	if nil != err {
		_t.Error(err)
	}
	_t.Log(rsp.Msg)
}
