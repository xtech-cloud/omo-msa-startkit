package subscriber

import (
	"context"

	"github.com/micro/go-micro/v2/logger"

	msa "omo-msa-startkit/proto/msa"
)

type StartKit struct{}

func (this *StartKit) Handle(_ctx context.Context, _msg *msa.Message) error {
	logger.Infof("Handler Received message: %v", _msg.Say)
	return nil
}

func Handler(_ctx context.Context, _msg *msa.Message) error {
	logger.Info("Function Received message: %v", _msg.Say)
	return nil
}
