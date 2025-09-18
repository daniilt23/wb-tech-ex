package main

import "fmt"

func worker(ch chan int, i int) {
	for val := range ch {
		fmt.Printf("Worker %d take data %d\n", i, val)
	}
}

func main() {
	var workerCount int
	if _, err := fmt.Scanf("%d", &workerCount); err != nil {
		fmt.Println(err)
		return
	}

	ch := make(chan int)
	for i := 0; i < workerCount; i++ {
		go worker(ch, i)
	}

	it := 0
	for {
		ch <- it
		it++
	}
}
