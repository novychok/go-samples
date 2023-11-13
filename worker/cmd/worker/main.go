package main

import (
	"fmt"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/novychok/go-samples/worker/internal"
)

func main() {
	if err := godotenv.Load("../../.env"); err != nil {
		fmt.Printf("error loading .env file %v\n", err)
	}

	db, err := internal.NewPostgres()
	if err != nil {
		fmt.Printf("error while running database %v\n", err)
	}

	storer := internal.NewStorer(db)
	handler := internal.NewHandler(storer)
	if err := storer.InitSchemas(); err != nil {
		fmt.Printf("error while initializing schemas %v\n", err)
	}

	http.HandleFunc("/callback", handler.Execute)

	if err := http.ListenAndServe(":9090", nil); err != nil {
		fmt.Printf("error while running the server %v\n", err)
	}
}
