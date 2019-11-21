package handler

import (
	"context"

	"github.com/micro/go-micro/util/log"

	shippydemo "github.com/alactic/shippydemo/proto/shippydemo"
)

type Shippydemo struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Shippydemo) Call(ctx context.Context, req *shippydemo.Request, rsp *shippydemo.Response) error {
	log.Log("Received Shippydemo.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *Shippydemo) Stream(ctx context.Context, req *shippydemo.StreamingRequest, stream shippydemo.Shippydemo_StreamStream) error {
	log.Logf("Received Shippydemo.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Logf("Responding: %d", i)
		if err := stream.Send(&shippydemo.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *Shippydemo) PingPong(ctx context.Context, stream shippydemo.Shippydemo_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Logf("Got ping %v", req.Stroke)
		if err := stream.Send(&shippydemo.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}
