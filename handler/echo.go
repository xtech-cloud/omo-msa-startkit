package handler

import (
	"context"
	"io"

	proto "omo-msa-startkit/proto/startkit"

	"github.com/micro/go-micro/v2/logger"
)

type Echo struct{}

// Call is a single request handler called via client.Call or the generated client code
func (this *Echo) Call(_ctx context.Context, _req *proto.Request, _rsp *proto.Response) error {
	logger.Infof("Received Echo.Call request: %v", _req)
	_rsp.Msg = _req.Name
	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (this *Echo) PingPong(_ctx context.Context, _stream proto.Echo_PingPongStream) error {
	for {
		req, err := _stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			return err
		}

		logger.Infof("Ping : %v", req.Stroke)
		if err = _stream.Send(&proto.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
	return nil
}
