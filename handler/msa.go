package handler

import (
	"context"

	"github.com/micro/go-micro/v2/util/log"

	msa "omo-msa-startkit/proto/msa"
)

type StartKit struct{}

// Call is a single request handler called via client.Call or the generated client code
func (this *StartKit) Call(_ctx context.Context, _req *msa.Request, _rsp *msa.Response) error {
	log.Log("Received StartKit.Call request")
	_rsp.Msg = _req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (this *StartKit) Stream(_ctx context.Context, _req *msa.StreamingRequest, _stream msa.StartKit_StreamStream) error {
	log.Logf("Received StartKit.Stream request with count: %d", _req.Count)

	for i := 0; i < int(_req.Count); i++ {
		log.Logf("Responding: %d", i)
		if err := _stream.Send(&msa.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (this *StartKit) PingPong(_ctx context.Context, _stream msa.StartKit_PingPongStream) error {
	for {
		req, err := _stream.Recv()
		if err != nil {
			return err
		}
		log.Logf("Got ping %v", req.Stroke)
		if err := _stream.Send(&msa.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
