package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(stopCh chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-stopCh:
			fmt.Println("end")
			return
		default:
			fmt.Println("goroutine work")
			time.Sleep(time.Nanosecond) // создаем задержку для процесса и меньшей нагрузки на CPU
		}
	}
}

func main() {
	stopCh := make(chan bool)

	var wg sync.WaitGroup

	wg.Add(1)
	go worker(stopCh, &wg)

	time.Sleep(time.Microsecond * 100) // ждем время до прерывания канала уведомителя
	stopCh <- true

	wg.Wait() // точно дожимаемся выполнения горутины
	fmt.Println("program finished")
}
