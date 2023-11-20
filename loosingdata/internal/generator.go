package internal

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"sync"
	"time"

	"github.com/novychok/go-samples/loosingdata/types"
)

var solt = "qwerty"

func GenerateData() []*types.Data {
	wg := sync.WaitGroup{}
	count := (5 + rand.Intn(5))

	dataSl := make([]*types.Data, 0, count)
	dataCh := make(chan *types.Data, count)
	tickerCh := make(chan time.Time)

	wg.Add(count)
	for i := 0; i < count; i++ {
		go func() {
			defer wg.Done()
			text := randText()
			hashText := createHash(text, solt)
			data := types.NewData(text, hashText)

			// Create fraud and truth text for Data
			go func() {
				for {
					randomTicker := time.NewTicker(500 + time.Duration(rand.Intn(1000))*time.Millisecond)
					t := <-randomTicker.C
					tickerCh <- t
				}
			}()
			select {
			case <-tickerCh:
				dataCh <- data
			case <-time.After(500 * time.Millisecond):
				data.Text = randText()
				dataCh <- data
			}
		}()
	}

	go func() {
		wg.Wait()
		close(dataCh)
	}()

	for v := range dataCh {
		dataSl = append(dataSl, v)
	}
	return dataSl
}

// RandText function generates random string using rand.Intn function
// to generate numbers from 65-90 and 97-122 to represend ASCII Table characters.
func randText() string {
	strLen := 10 + rand.Intn(19)
	letsl := make([]rune, strLen)

	for i := 0; i < strLen; i++ {
		uppercase := (int32)(65 + rand.Intn(25))
		lowercase := (int32)(97 + rand.Intn(25))
		upOrNot := rand.Intn(2)

		if upOrNot == 1 {
			letsl = append(letsl, uppercase)
			continue
		}
		letsl = append(letsl, lowercase)
	}
	return string(letsl)
}

func createHash(text, solt string) string {
	textHash := md5.Sum([]byte(text + solt))
	return hex.EncodeToString(textHash[:])
}
