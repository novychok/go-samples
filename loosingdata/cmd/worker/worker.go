package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/novychok/go-samples/loosingdata/internal"
)

func main() {
	listenAddr := ":8100"

	http.HandleFunc("/data", internal.HandleLoosingData)

	go func() { log.Fatal(http.ListenAndServe(listenAddr, nil)) }()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt)
	<-sigCh

	fmt.Println("interrupt worker")
}
