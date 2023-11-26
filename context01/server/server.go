package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	wg := &sync.WaitGroup{}
	urlsch := make(chan string)

	urls := []string{
		"https://google.com",
		"https://google.com",
		"https://google.com",
	}

	for _, url := range urls {
		wg.Add(1)
		go fetchAPI(ctx, url, urlsch, wg)
	}

	go func() {
		wg.Wait()
		close(urlsch)
	}()

	for v := range urlsch {
		fmt.Println(v)
	}

}

func fetchAPI(ctx context.Context, url string, urslch chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
		return
	}
	// time.Sleep(3 * time.Second)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer resp.Body.Close()

	urslch <- fmt.Sprintf("Response from: %s %d", url, resp.StatusCode)
}
