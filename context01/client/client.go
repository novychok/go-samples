package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithDeadlineCause(context.TODO(), time.Now().Add(10*time.Millisecond),
		fmt.Errorf("errorik"))
	defer cancel()

	// Simulate work
	time.Sleep(200 * time.Millisecond)

	// Print the error cause
	fmt.Println(ctx.Err()) // prints "context deadline exceeded: RPC timeout"
}
