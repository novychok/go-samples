package main

import (
	"flag"
	"fmt"

	"github.com/novychok/go-samples/realtime/internal/handler/bookapi"
	natsHandler "github.com/novychok/go-samples/realtime/internal/handler/nats"
	"github.com/novychok/go-samples/realtime/internal/handler/websocketapi"
	log "github.com/novychok/go-samples/realtime/internal/pkg/log"
	nats "github.com/novychok/go-samples/realtime/internal/pkg/nats"

	"github.com/novychok/go-samples/realtime/internal/pkg/server"
	"github.com/novychok/go-samples/realtime/internal/service/book"
	"github.com/novychok/go-samples/realtime/internal/service/realtime"
)

func main() {

	port := flag.String("port", "3380", "bookapi server listenAddr")
	flag.Parse()

	slogger := log.New()

	natsClient, cleanup, err := nats.New()
	if err != nil {
		slogger.Error(fmt.Sprintf("error to start nats client: %s", err.Error()))
	}
	defer cleanup()

	bookService := book.New(slogger, natsClient)
	realtimeService := realtime.New()

	bookapiHandler := bookapi.New(bookService)
	websocketapiHandler := websocketapi.New(realtimeService)
	natsHnadler := natsHandler.New(realtimeService)

	server := server.New(*port, *bookapiHandler, *natsHnadler, *websocketapiHandler, natsClient)
	slogger.Error(fmt.Sprintf("error to start the server on port - :%s",
		*port), server.Start())
}
