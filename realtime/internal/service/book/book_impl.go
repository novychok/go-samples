package book

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/nats-io/nats.go/jetstream"
	"github.com/novychok/go-samples/realtime/internal/entity"
	"github.com/novychok/go-samples/realtime/internal/repository"
	"github.com/novychok/go-samples/realtime/internal/service"
)

type srv struct {
	l          *slog.Logger
	natsClient jetstream.JetStream
	repo       repository.Book
}

func (s *srv) Create(ctx context.Context, book *entity.Book) (int, error) {

	bookId, err := s.repo.Create(ctx, book)
	if err != nil {
		return 0, err
	}

	_, err = s.natsClient.Publish(ctx, "collections.update", []byte(fmt.Sprintf("%d", bookId)))
	if err != nil {
		s.l.ErrorContext(ctx, "failed to publish a message", "err", err)
		return 0, err
	}

	return bookId, nil
}

func New(l *slog.Logger, natsClient jetstream.JetStream, repo repository.Book) service.Books {
	return &srv{
		l:          l,
		natsClient: natsClient,
		repo:       repo,
	}
}
