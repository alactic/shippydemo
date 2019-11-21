package subscriber

import (
	"context"
	"github.com/micro/go-micro/util/log"

	shippydemo "github.com/alactic/shippydemo/proto/shippydemo"
)

type Shippydemo struct{}

func (e *Shippydemo) Handle(ctx context.Context, msg *shippydemo.Message) error {
	log.Log("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *shippydemo.Message) error {
	log.Log("Function Received message: ", msg.Say)
	return nil
}
