package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
)

// для завершения работы всех горутин выбрал контекст, так как
// go-каналы не работают широковещательно, только одна горутина получит один os.Signal.
// также нет никакой гарантии, какая именно горутина его получит.
// контекст автоматически уведомляет всех воркеров

func worker(ctx context.Context, in chan int, index int) {
	for val := range in {
		select {
		case <-ctx.Done():
			fmt.Printf("stopped %d worker\n", index)
			return
		default:
			fmt.Printf("%d\n", val)
		}
	}
}

func main() {
	var workerCount int
	if _, err := fmt.Scanf("%d", &workerCount); err != nil {
		fmt.Println(err)
		return
	}
	ctx, cancel := context.WithCancel(context.Background())
	ch := make(chan int)

	go func() {
		exit := make(chan os.Signal, 1)
		signal.Notify(exit, os.Interrupt)
		<-exit
		cancel()
	}()

	for i := 0; i < workerCount; i++ {
		go worker(ctx, ch, i)
	}

	it := 0
	for {
		ch <- it
		it++
	}
}
