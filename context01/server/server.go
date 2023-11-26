package main

import (
	"context"
	"fmt"
	"math"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 6*time.Second)
	defer cancel()

	doStuff2(ctx)
}

func doStuff2(ctx context.Context) {
	timer := time.Now()
	select {
	case <-ctx.Done():
		fmt.Println("Time is up: ", math.Round(time.Since(timer).Seconds()))
		return
	case <-time.After(8 * time.Second):
		fmt.Println("The task is complited: ", math.Round(time.Since(timer).Seconds()))
	}
}
