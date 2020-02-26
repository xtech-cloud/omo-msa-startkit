package subscriber

import (
	"context"

	"github.com/micro/go-micro/v2/util/log"

	msa "omo-msa-startkit/proto/msa"
)

type StartKit struct{}

func (this *StartKit) Handle(_ctx context.Context, _msg *msa.Message) error {
	log.Log("Handler Received message: ", _msg.Say)
	return nil
}

func Handler(_ctx context.Context, _msg *msa.Message) error {
	log.Log("Function Received message: ", _msg.Say)
	return nil
}
