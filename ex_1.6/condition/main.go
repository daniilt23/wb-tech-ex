package main

import (
	"fmt"
	"time"
)

func worker(timer <-chan time.Time) {
	for {
		select {
		case <-timer:
			fmt.Println("end")
			return
		default:
			fmt.Println("goroutine run")
		}
	}
}

func main() {
	timer := time.After(time.Second)

	go worker(timer)

	time.Sleep(time.Second)
}
