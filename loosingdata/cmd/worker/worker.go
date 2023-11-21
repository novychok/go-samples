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

	db, err := internal.SetupBbolt()
	if err != nil {
		fmt.Println(err)
		return
	}

	repo := internal.NewBolt(db)
	handler := internal.NewHandler(repo)

	http.HandleFunc("/data", handler.HandleGetData)
	http.HandleFunc("/fraud_data", handler.HandleGetFraudData)
	http.HandleFunc("/fraud_data/:id", handler.HandleGetFraudDataByID)

	go func() { log.Fatal(http.ListenAndServe(listenAddr, nil)) }()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt)
	<-sigCh

	fmt.Println("interrupt worker")
}
