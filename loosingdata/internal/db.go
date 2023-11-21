package internal

import (
	"encoding/json"
	"fmt"

	"github.com/novychok/go-samples/loosingdata/types"
	"go.etcd.io/bbolt"
)

type Bolt struct {
	db *bbolt.DB
}

func NewBolt(db *bbolt.DB) *Bolt {
	return &Bolt{db: db}
}

func (b *Bolt) SetupBbolt() (*bbolt.DB, error) {
	db, err := bbolt.Open("../db/data.db", 0600, nil)
	if err != nil {
		return nil, fmt.Errorf("error while connecting to db: %v", err)
	}
	err = db.Update(func(tx *bbolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("DATA_DB"))
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("error while creating buckets: %v", err)
	}

	return db, nil
}

func (b *Bolt) sendData(db *bbolt.DB, data *types.Data) error {
	dataBytes, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("error while marshaling data: %v", err)
	}
	err = db.Update(func(tx *bbolt.Tx) error {
		err := tx.Bucket([]byte("DATA_DB")).Put([]byte("DATA"), dataBytes)
		if err != nil {
			return fmt.Errorf("error while execute transaction: %v", err)
		}

		return nil
	})
	if err != nil {
		return fmt.Errorf("error while update DATA_DB: %v", err)
	}

	return nil
}

func (b *Bolt) retrieveData(db *bbolt.DB, path []byte) error {
	err := db.View(func(tx *bbolt.Tx) error {
		data := tx.Bucket([]byte("DATA_DB")).Get(path)
		// Make to golang structure
		fmt.Printf("FRAUD DATA: %+v\n", data)
		return nil
	})
	if err != nil {
		return fmt.Errorf("error while retrieve DATA from DATA_DB: %v", err)
	}

	return nil
}
