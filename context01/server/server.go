package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if err := userRegistration(ctx); err != nil {
		fmt.Println(err)
	}
}

func userRegistration(ctx context.Context) error {
	id := 1 + rand.Intn(5)
	ctx = context.WithValue(ctx, "userId", id)

	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
	defer cancel()

	select {
	case <-ctx.Done():
		return fmt.Errorf("time is up for registration: %v", ctx.Err())
	case <-time.After(1 * time.Second):
		msg := fmt.Sprintf("registration for userId[%d] is done, proceed user login\n", id)
		if err := userLogin(ctx, msg); err != nil {
			return err
		}
		return nil
	}

}

func userLogin(ctx context.Context, msg string) error {
	getCtxID, ok := ctx.Value("userId").(int)
	if !ok {
		fmt.Printf("wrong custing?\n")
		return nil
	}

	select {
	case <-ctx.Done():
		return fmt.Errorf("time is up for userId[%d] login", getCtxID)
	case <-time.After(2 * time.Second):
		fmt.Print(msg)
		return nil
	}
}
