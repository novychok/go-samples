package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

type userId string

func main() {
	ctx, cancel := context.WithCancel(nil)
	defer cancel()

	if err := userRegistration(ctx); err != nil {
		fmt.Println(err)
	}
}

func userRegistration(ctx context.Context) error {
	id := 1 + rand.Intn(5)

	// userId = "userId"

	ctx, cancel := context.WithTimeout(context.WithValue(ctx, "userId", id), 2*time.Second)
	defer cancel()

	select {
	case <-ctx.Done():
		return fmt.Errorf("time is up for registration: %v", ctx.Err())
	case <-time.After(3 * time.Second):
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
