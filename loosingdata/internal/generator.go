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
			msg := randMsg()
			msgHash := createHash(msg, solt)
			data := types.NewData(msg, msgHash)

			// Create fraud and truth Data
			go func() {
				for {
					randomTicker := time.NewTicker(500 + time.Duration(rand.Intn(500))*time.Millisecond)
					t := <-randomTicker.C
					tickerCh <- t
				}
			}()
			select {
			case <-tickerCh:
				dataCh <- data
			case <-time.After(300 * time.Millisecond):
				data.Message = randMsg()
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

// randMsg function generates random string using rand.Intn function
// to generate numbers from 65-90 and 97-122 to represend ASCII Table characters.
func randMsg() string {
	msgLen := 10 + rand.Intn(19)
	msgSl := make([]rune, 0, msgLen)

	for i := 0; i < msgLen; i++ {
		uppercase := (int32)(65 + rand.Intn(25))
		lowercase := (int32)(97 + rand.Intn(25))
		upOrNot := rand.Intn(2)

		if upOrNot == 1 {
			msgSl = append(msgSl, uppercase)
			continue
		}
		msgSl = append(msgSl, lowercase)
	}
	return string(msgSl)
}

// createHash creates hash from message + solt strings
func createHash(msg, solt string) string {
	msgHash := md5.Sum([]byte(msg + solt))
	return hex.EncodeToString(msgHash[:])
}
