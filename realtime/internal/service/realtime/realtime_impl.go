package realtime

import (
	"context"

	"github.com/novychok/go-samples/realtime/internal/repository"
	"github.com/novychok/go-samples/realtime/internal/service"
)

type srv struct {
	broadcast  chan string
	repository repository.Book
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

		// event, err := s.repository.GetEvent(ctx)
		// if err != nil {
		// 	// TODO: slog
		// 	log.Println(err)
		// 	continue
		// }

		// if event.ID == 0 {
		// 	// TODO: slog
		// 	log.Println(err)
		// 	continue
		// }

		// if err := s.repository.SetDone(ctx, event.ID); err != nil {
		// 	// TODO: slog
		// 	log.Println(err)
		// 	continue
		// }
	}
}

func New(repository repository.Book) service.Realtime {
	return &srv{
		repository: repository,
		broadcast:  make(chan string, 1024),
	}
}
