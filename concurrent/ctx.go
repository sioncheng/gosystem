package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx,cancel := context.WithCancel(context.Background())
	b := false
	for i := 0; !b ; i++ {
		select {
		case <- ctx.Done():
			b = true
		case <- time.After(time.Second):
			fmt.Println("tick", i)
			if 5 == i {
				cancel()
			}
		}
	}

	fmt.Println("...")

	ctx, cancel = context.WithDeadline(context.Background(), time.Now().Add(5 * time.Second))
	b = false
	for i := 0; !b ; i++ {
		select {
		case <- ctx.Done():
			fmt.Println("exit", ctx.Err())
			b = true
		case <- time.After(time.Second):
			fmt.Println("tick", i)
			if 10 == i {
				cancel()
			}
		}
	}
}
