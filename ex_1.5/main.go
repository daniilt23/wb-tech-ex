package main

import (
	"fmt"
	"time"
)

func writer(timer <-chan time.Time, ch chan int) {
	it := 0
	for {
		select {
		case <-timer:
			fmt.Println("time is out")
			close(ch)
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

	timer := time.After(time.Duration(N) * time.Second)

	ch := make(chan int)
	go writer(timer, ch)

	for val := range ch {
		fmt.Println(val)
	}
}
