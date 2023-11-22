package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/joho/godotenv"
	"github.com/novychok/go-samples/loosingdata/internal"
)

func main() {
	if err := godotenv.Load(); err != nil {
		fmt.Printf("error to load .env file: %v\n", err)
	}

	if err := internal.HandleSendData(); err != nil {
		fmt.Printf("error to send Data: %v\n", err)
	}

	go func() { log.Fatal(http.ListenAndServe(":"+os.Getenv("MASTER_ADDR"), nil)) }()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt)
	<-sigCh

	fmt.Println("interrupt master")
}
