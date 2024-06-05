package realtime

import (
	"context"

	"github.com/novychok/go-samples/realtime/internal/service"
)

type srv struct {
	broadcast chan string
}

func (s *srv) PublishMessage(ctx context.Context, message string) error {
	select {
	case s.broadcast <- message:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

func (s *srv) SubscribeToMessages(ctx context.Context, handler func(message string)) error {
	for {
		select {
		case msg := <-s.broadcast:
			handler(msg)
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}

func New() service.Realtime {
	return &srv{
		broadcast: make(chan string, 1024),
	}
}
