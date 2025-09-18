package main

import (
	"fmt"
	"time"
)

func worker(stop_ch chan bool) {
	for {
		select {
		case <-stop_ch:
			fmt.Println("end")
			return
		default:
			fmt.Println("goroutine work")
		}
	}
}

func main() {
	stop_ch := make(chan bool)

	go worker(stop_ch)

	go func() {
		time.Sleep(time.Microsecond * 100) // ждем время до прерывания канала уведомителя
		stop_ch <- true
	}()

	time.Sleep(time.Millisecond) // даем время на завершение горутин
}
