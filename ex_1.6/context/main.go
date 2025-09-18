package main

import (
	"context"
	"fmt"
	"time"
)

func worker(ctx context.Context, ch chan bool) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("end")
			close(ch)
			return
		default:
			fmt.Println("goroutine work")
		}
	}
}

func main() {
	// создадим контекст который завершается через одну секунду
	ctx, stop := context.WithTimeout(context.Background(), time.Second)
	defer stop()

	ch := make(chan bool)

	go worker(ctx, ch)

	for val := range ch {
		fmt.Println(val)
	}
}
