package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(ch chan int, w int, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range ch {
		fmt.Printf("worker %d starter job %d\n", w, job)
		time.Sleep(time.Second)
		fmt.Printf("worker %d complete job %d\n", w, job)
	}
}

func main() {
	t1 := time.Now()
	numJobs := 30
	numWorkers := 30
	jobs := make([]int, 0, numJobs)
	for i := range numJobs {
		jobs = append(jobs, i)
	}

	ch := make(chan int)
	go func() {
		for _, val := range jobs {
			ch <- val
		}
		close(ch)
	}()

	wg := sync.WaitGroup{}
	wg.Add(numWorkers)
	for i := 1; i <= numWorkers; i++ {
		go worker(ch, i, &wg)
	}
	wg.Wait()
	fmt.Printf("%v", time.Since(t1))
}
