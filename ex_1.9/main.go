package main

import (
	"fmt"
	"sync"
)

func generate(in chan<- int, mas [10]int, wg *sync.WaitGroup) {
	defer wg.Done()
	defer close(in)
	for _, val := range mas {
		in <- val
	}
}

func calculate(in <-chan int, out chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	defer close(out)
	for val := range in {
		out <- val * 2
	}
}

func main() {
	in := make(chan int)
	out := make(chan int)
	mas := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} // про размер и числа ничего не сказано создадим массив размером 10

	wg := sync.WaitGroup{}
	wg.Add(2)
	go generate(in, mas, &wg)
	go calculate(in, out, &wg)

	for val := range out {
		fmt.Println(val)
	}
	wg.Wait()

	fmt.Println("Pipiline completed")
}
