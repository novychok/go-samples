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

	db, err := internal.SetupBbolt()
	if err != nil {
		fmt.Printf("%+v\n", err)
		return
	}

	repo := internal.NewBbolt(db)
	handler := internal.NewHandler(repo)

	http.HandleFunc("/data", handler.HandleGetData)
	http.HandleFunc("/all_fraud_data", handler.HandleGetFraudData)
	http.HandleFunc("/fraud_data", handler.HandleGetFraudDataByID)

	go func() { log.Fatal(http.ListenAndServe(":"+os.Getenv("WORKER_ADDR"), nil)) }()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt)
	<-sigCh

	fmt.Println("interrupt worker")
}
