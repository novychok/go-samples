package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/novychok/go-samples/loosingdata/internal"
)

func main() {
	listenAddr := ":8090"
	client := &http.Client{Timeout: 1 * time.Second}

	go func() {
		for {
			time.Sleep(5 * time.Second)

			dataSl := internal.GenerateData()
			jsonData, err := json.Marshal(dataSl)
			if err != nil {
				fmt.Println(err)
				continue
			}

			body := bytes.NewBuffer(jsonData)
			resp, err := client.Post("http://127.0.0.1:8100/data", "application/json", body)
			if err != nil {
				fmt.Println(err)
				continue
			}
			resp.Body.Close()
		}
	}()

	go func() { log.Fatal(http.ListenAndServe(listenAddr, nil)) }()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt)
	<-sigCh

	fmt.Println("interrupt master")
}
