package book

import (
	"context"
	"log/slog"

	"github.com/nats-io/nats.go/jetstream"
	"github.com/novychok/go-samples/realtime/internal/entity"
	"github.com/novychok/go-samples/realtime/internal/service"
)

type srv struct {
	l          *slog.Logger
	natsClient jetstream.JetStream
}

func (s *srv) Create(ctx context.Context, book *entity.Book) (*entity.Book, error) {

	byteCollectionDataID := book.ID
	_, err := s.natsClient.Publish(ctx, "collections.update", []byte(byteCollectionDataID))
	if err != nil {
		s.l.ErrorContext(ctx, "failed to publish a message", "err", err)
		return nil, err
	}

	return book, nil
}

func (s *srv) FindAll(ctx context.Context) ([]*entity.Book, error) {
	panic("implement me")
}

func (s *srv) GetByID(ctx context.Context, id string) (*entity.Book, error) {
	panic("implement me")
}

func New(l *slog.Logger, natsClient jetstream.JetStream) service.Books {
	return &srv{
		l:          l,
		natsClient: natsClient,
	}
}
