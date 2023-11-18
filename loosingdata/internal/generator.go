package internal

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"sync"

	"github.com/novychok/go-samples/loosingdata/types"
)

func GenerateData() []*types.Data {
	wg := sync.WaitGroup{}
	count := (5 + rand.Intn(5))

	dataSl := make([]*types.Data, 0, count)
	dataCh := make(chan *types.Data, count)

	wg.Add(count)
	for i := 0; i < count; i++ {
		go func() {
			defer wg.Done()
			text := randText()
			data := types.NewData(text)

			dataCh <- data
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

func Encode(text string) string {
	hash := md5.Sum([]byte(text))
	return string(hash[:])
}

func Decode(text string) string {
	hash := hex.EncodeToString([]byte(text))
	return hash
}
