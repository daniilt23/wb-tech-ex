package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(stop_ch chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-stop_ch:
			fmt.Println("end")
			return
		default:
			fmt.Println("goroutine work")
			time.Sleep(time.Nanosecond) // создаем задержку для процесса и меньшей нагрузки на CPU
		}
	}
}

func main() {
	stop_ch := make(chan bool)

	var wg sync.WaitGroup

	wg.Add(1)
	go worker(stop_ch, &wg)

	time.Sleep(time.Microsecond * 100) // ждем время до прерывания канала уведомителя
	stop_ch <- true

	wg.Wait() // точно дожимаемся выполнения горутины
	fmt.Println("program finished")
}
