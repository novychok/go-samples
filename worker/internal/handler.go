package internal

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	"github.com/novychok/go-samples/worker/types"
)

type Handler struct {
	storer *Storer
}

func NewHandler(storer *Storer) *Handler {
	return &Handler{storer: storer}
}

func (h *Handler) Execute(w http.ResponseWriter, r *http.Request) {
	objects, err := getObjectIds(r)
	if err != nil {
		return
	}

	wg := &sync.WaitGroup{}
	objectAmount := len(objects.ID)
	objectChan := make(chan types.Object, objectAmount)

	for i := range objects.ID {
		wg.Add(1)
		objID := objects.ID[i]

		go func() {
			defer wg.Done()

			object, err := getObject(objID)
			if err != nil {
				return
			}

			// fmt.Println(object)
			objectChan <- object
		}()

	}

	go func() {
		wg.Wait()
		close(objectChan)
	}()

	var objectOnline int
	for v := range objectChan {
		if !v.Online {
			continue
		}

		if err := h.storer.StoreObject(v); err != nil {
			fmt.Println(err)
			return
		}
		objectOnline++
	}

	fmt.Printf("objects: [amount %d] [online %d] [saved %d]\n", objectAmount, objectOnline, objectOnline)
}

func getObject(objID int) (types.Object, error) {

	url := fmt.Sprintf("http://localhost:9010/objects/%d", objID)
	resp, err := http.Get(url)
	if err != nil {
		return types.Object{}, fmt.Errorf("failed to get object with id [%d] | err: %s", objID, err)
	}
	defer resp.Body.Close()

	var obj types.Object
	if err := json.NewDecoder(resp.Body).Decode(&obj); err != nil {
		return types.Object{}, fmt.Errorf("failed to decode the response body | err: %s", err)
	}
	return obj, nil
}

func getObjectIds(r *http.Request) (*types.ObjectsIDs, error) {
	var err error

	var objIds types.ObjectsIDs
	if err = json.NewDecoder(r.Body).Decode(&objIds); err != nil {
		return nil, fmt.Errorf("failed to decode object ids | err: %s", err)
	}
	if len(objIds.ID) == 0 {
		return nil, fmt.Errorf("get [%d] id's from objectIds: %s", len(objIds.ID), err)
	}

	return &objIds, nil
}
