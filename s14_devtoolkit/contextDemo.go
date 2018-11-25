package main

import (
	"context"
	"fmt"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	ch := generate(ctx)
	for n := range ch {
		fmt.Println(n)
		if n == 5 {
			cancel()
			break
		}
	}
}

func generate(ctx context.Context) <-chan int {
	ch := make(chan int)

	go func() {
		var n int
		for {
			select {
			case <-ctx.Done():
				return
			case ch <- n:
				n++
			}
		}
	}()

	return ch
}
