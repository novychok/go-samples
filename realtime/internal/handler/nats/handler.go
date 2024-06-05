package nats

import (
	"context"
	"log"

	"github.com/novychok/go-samples/realtime/internal/service"
)

type Handler struct {
	realtimeService service.Realtime
}

func (h *Handler) CollectionsWebHook(ctx context.Context, data []byte) error {
	message := string(data)

	if err := h.realtimeService.PublishMessage(ctx, message); err != nil {
		log.Println("Failed to publish message:", err)
		return err
	}

	return nil
}

func New(realtimeService service.Realtime) *Handler {
	return &Handler{
		realtimeService: realtimeService,
	}
}
