package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func worker(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("end")
			return
		default:
			time.Sleep(time.Nanosecond) //создаем задержку для имитации работы и облегчения нагрузки
			fmt.Println("goroutine work")
		}
	}
}

func main() {
	// создадим контекст который завершается через одну секунду
	ctx, stop := context.WithTimeout(context.Background(), time.Second)
	defer stop()

	var wg sync.WaitGroup

	wg.Add(1)
	go worker(ctx, &wg)
	wg.Wait() //дожидаемся в главной горутине звершения воркера

	fmt.Println("program finished")
}
