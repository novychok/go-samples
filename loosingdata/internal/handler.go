package internal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/novychok/go-samples/loosingdata/types"
)

type Handler struct {
	repo *Bolt
}

func NewHandler(repo *Bolt) *Handler {
	return &Handler{repo: repo}
}

func HandleSendData() error {
	client := &http.Client{Timeout: 1 * time.Second}
	go func() {
		for {
			time.Sleep(5 * time.Second)

			dataSl := GenerateData()

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
	return nil
}

func (h *Handler) HandleGetData(w http.ResponseWriter, r *http.Request) {
	var data []*types.Data
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		fmt.Printf("error to decode data: %v", err)
		return
	}
	for _, v := range data {
		hash := createHash(v.Text, solt)
		if hash != v.TextHash {
			// Function to save fraud data to bbolt
			h.repo.saveFraudData(v)
			continue
		}
	}
}

func (h *Handler) HandleGetFraudData(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) HandleGetFraudDataByID(w http.ResponseWriter, r *http.Request) {

}
