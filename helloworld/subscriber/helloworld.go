package subscriber

import (
	"context"

	log "github.com/ebelanja/go-micro/logger"

	helloworld "github.com/ebelanja/micro-examples/helloworld/proto"
)

type Helloworld struct{}

func (e *Helloworld) Handle(ctx context.Context, msg *helloworld.Message) error {
	log.Info("Handler Received message: ", msg.Say)
	return nil
}

func Handler(ctx context.Context, msg *helloworld.Message) error {
	log.Info("Function Received message: ", msg.Say)
	return nil
}
