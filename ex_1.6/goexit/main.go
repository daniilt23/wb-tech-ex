package main

import (
	"fmt"
	"runtime"
	"sync"
)

func worker(i int, wg *sync.WaitGroup) {
	defer wg.Done()
	defer fmt.Printf("worker %d was killed\n", i)

	for it := range 5 { // имитируем небольшую работу воркера
		if it == 3 { // создаем условие для вызова Goexit()
			runtime.Goexit() // убиваем воркер, без влияния на остальные воркеры
		}

		fmt.Printf("goroutine %d is worked on process %d\n", i, it)
	}
}

func main() {
	var wg sync.WaitGroup

	for i := 1; i < 3; i++ { // создаем два воркера чтобы показать прерывания каждого
		wg.Add(1)
		go worker(i, &wg)
	}
	wg.Wait()

	fmt.Println("program finished")
}
