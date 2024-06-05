package nats

import (
	"fmt"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
)

var natsUri = "nats://nats-container:4222"

func New() (jetstream.JetStream, func(), error) {
	nc, err := nats.Connect(natsUri)
	if err != nil {
		return nil, nil, fmt.Errorf("error to connect conn: [%s]", err.Error())
	}

	js, err := jetstream.New(nc)
	if err != nil {
		return nil, nil, fmt.Errorf("error to connect jetstream: [%s]", err.Error())
	}

	cleanup := func() {
		nc.Close()
	}

	return js, cleanup, nil
}
