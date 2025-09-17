package main

import (
	"context"
	"fmt"
	"time"
)

func writer(ctx context.Context, ch chan int) {
	it := 0
	for {
		select {
		case <-ctx.Done():
			fmt.Println("time is out")
			return
		default:
			ch <- it
			it++
		}
	}
}

func main() {
	var N int
	if _, err := fmt.Scanf("%d", &N); err != nil {
		fmt.Println(err)
		return
	}

	//самым оптимальным будет просто завершать контекст по времени ввода
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(N)*time.Second)
	defer cancel()

	ch := make(chan int)
	go writer(ctx, ch)

	for {
		select {
		case <-ctx.Done():
			return
		default:
			fmt.Println(<-ch)
		}
	}
}
