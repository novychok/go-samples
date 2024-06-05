package websocketapi

import (
	"context"
	"log"
	"net/http"
	"sync"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/novychok/go-samples/realtime/internal/service"
)

type Handler struct {
	mu sync.Mutex

	connections     map[string]*websocket.Conn
	realtimeService service.Realtime
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (h *Handler) WebscocketHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	connectionID := uuid.New().String()

	h.mu.Lock()
	h.connections[connectionID] = conn
	h.mu.Unlock()

	defer func() {
		h.mu.Lock()
		delete(h.connections, connectionID)
		h.mu.Unlock()
	}()

	for {
		_, _, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			break
		}
	}
}

func (h *Handler) broadcastMessages() error {

	ctx := context.Background() // todo use app context
	err := h.realtimeService.SubscribeToMessages(ctx, func(message string) {
		h.mu.Lock()
		defer h.mu.Unlock()

		for _, conn := range h.connections {
			if err := conn.WriteMessage(websocket.TextMessage, []byte(message)); err != nil {
				log.Println(err)
			}
		}
	})

	return err
}

func New(realtimeService service.Realtime) *Handler {
	handler := &Handler{
		connections:     make(map[string]*websocket.Conn),
		realtimeService: realtimeService,
	}

	go func() {
		if err := handler.broadcastMessages(); err != nil {
			log.Println(err)
		}
	}() // todo handle error and call it as handler listener

	return handler
}
