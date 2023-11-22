package internal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/novychok/go-samples/loosingdata/types"
	"go.etcd.io/bbolt"
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
		hash := createHash(v.Message, solt)
		if hash != v.MessageHash {
			// Function to save fraud data to bbolt
			h.repo.saveFraudData(v)
			continue
		}
	}
}

func (h *Handler) HandleGetFraudData(w http.ResponseWriter, r *http.Request) {
	godotenv.Load()
	err := h.repo.db.View(func(tx *bbolt.Tx) error {
		tx.Bucket([]byte(os.Getenv("DB_NAME"))).ForEach(func(k, v []byte) error {
			var data types.Data
			reader := bytes.NewReader(v)
			if err := json.NewDecoder(reader).Decode(&data); err != nil {
				return fmt.Errorf("error while decode data: %v", err)
			}
			bb, err := json.MarshalIndent(data, "", "     ")
			if err != nil {
				return err
			}
			w.Write(bb)
			return nil
		})
		return nil
	})
	if err != nil {
		fmt.Printf("error while update DATA_DB: %v", err)
		return
	}
}

func (h *Handler) HandleGetFraudDataByID(w http.ResponseWriter, r *http.Request) {
	godotenv.Load()
	id := r.URL.Query().Get("id")
	var dataBytes []byte
	var data *types.Data
	err := h.repo.db.View(func(tx *bbolt.Tx) error {
		dataBytes = tx.Bucket([]byte(os.Getenv("DB_NAME"))).Get([]byte(id))
		return nil
	})
	if err != nil {
		fmt.Printf("error while try to View data: %v", err)
		return
	}

	bufferedData := bytes.NewBuffer(dataBytes)
	if err := json.NewDecoder(bufferedData).Decode(&data); err != nil {
		fmt.Printf("error while decode Data bytes to Data structure: %v", err)
	}

	jsonData, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		fmt.Printf("error while marshal Data to json: %v", err)
		return
	}
	w.Write(jsonData)
}
