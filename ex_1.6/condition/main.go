package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(timer <-chan time.Time, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-timer:
			fmt.Println("end")
			return
		default:
			time.Sleep(time.Millisecond) //создаем задержку для имитации работы и облегчения нагрузки
			fmt.Println("goroutine run")
		}
	}
}

func main() {
	// сделаем что наше условие это время
	timer := time.After(time.Second)
	var wg sync.WaitGroup

	wg.Add(1)
	go worker(timer, &wg)
	wg.Wait() //дожидаемся в главной горутине звершения воркера

	fmt.Println("program finished")
}
