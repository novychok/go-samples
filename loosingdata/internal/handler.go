package internal

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/novychok/go-samples/loosingdata/types"
)

func HandleLoosingData(w http.ResponseWriter, r *http.Request) {
	var data []*types.Data
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		fmt.Printf("error to decode data: %v", err)
		return
	}

	for _, v := range data {
		fmt.Println("worker:::: ", v)
	}
}
