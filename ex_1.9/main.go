package main

import (
	"fmt"
	"sync"
)

func getValue(in chan<- int, wg *sync.WaitGroup, mas [10]int) {
	defer wg.Done()
	defer close(in)
	for _, val := range mas {
		in <- val
	}
}

func calculateResult(result chan<- int, in <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	defer close(result)
	for val := range in {
		result <- val
	}
}

func main() {
	mas := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	in := make(chan int)
	result := make(chan int)
	wg := sync.WaitGroup{}

	wg.Add(2)
	go getValue(in, &wg, mas)
	go calculateResult(result, in, &wg)
	go func() {
		wg.Wait()
	}()

	for val := range result {
		fmt.Println(val)
	}
}
