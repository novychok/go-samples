package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"

	"github.com/nats-io/nats.go/jetstream"
	"github.com/novychok/go-samples/realtime/internal/handler/bookapi"
	"github.com/novychok/go-samples/realtime/internal/handler/websocketapi"

	natsHandler "github.com/novychok/go-samples/realtime/internal/handler/nats"
)

type Server struct {
	httpServer *http.Server

	listenAddr          string
	bookApiHandler      bookapi.Handler
	websocketApiHandler websocketapi.Handler
	natsHandler         natsHandler.Handler

	natsClient jetstream.JetStream

	mu            sync.Mutex
	consumernames map[string]bool
}

func (s *Server) Start() error {

	router := http.NewServeMux()

	s.httpServer = &http.Server{
		Addr:    fmt.Sprintf(":%s", s.listenAddr),
		Handler: router,
	}

	router.HandleFunc("POST /api/v1/books", s.bookApiHandler.CreateBook)

	router.HandleFunc("GET /api/v1/ws/", s.websocketApiHandler.WebscocketHandler)

	if err := s.NatsStart(); err != nil {
		log.Println(err)
	}

	fmt.Printf("server started on - :%s\n", s.listenAddr)
	return s.httpServer.ListenAndServe()
}

func (s *Server) NatsStart() error {

	ctx := context.Background()
	stream, err := s.natsClient.CreateOrUpdateStream(ctx, jetstream.StreamConfig{
		Name:      "collections",
		Subjects:  []string{"collections.*"},
		Retention: jetstream.InterestPolicy,
	})
	if err != nil {
		return err
	}

	hostname, err := os.Hostname()
	if err != nil {
		log.Println(err)
	}

	fmt.Println(hostname)

	consumer, err := stream.CreateOrUpdateConsumer(ctx, jetstream.ConsumerConfig{
		Durable:   hostname,
		AckPolicy: jetstream.AckExplicitPolicy,
	})
	if err != nil {
		return err
	}

	_, err = consumer.Consume(func(msg jetstream.Msg) {
		switch msg.Subject() {
		case "collections.update":
			if err := s.natsHandler.CollectionsWebHook(ctx, msg.Data()); err != nil {
				log.Println(err)
			}
		}
		if err := msg.Ack(); err != nil {
			fmt.Println(err)
		}
	})
	if err != nil {
		return err
	}

	return nil
}

func New(listenAddr string, bookApiHandler bookapi.Handler, natsHandler natsHandler.Handler,
	websocketApiHandler websocketapi.Handler, natsClient jetstream.JetStream) *Server {
	return &Server{
		listenAddr: listenAddr,

		bookApiHandler:      bookApiHandler,
		websocketApiHandler: websocketApiHandler,
		natsHandler:         natsHandler,

		natsClient: natsClient,

		consumernames: make(map[string]bool),
	}
}
