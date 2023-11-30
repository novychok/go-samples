package main

import (
	"log"
	"os"
	"os/signal"
)

func main() {
	tcpTransport := NewTCPTransport(":8080")

	log.Fatal(tcpTransport.ListenAndAccept())

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)
	<-sigChan
}
