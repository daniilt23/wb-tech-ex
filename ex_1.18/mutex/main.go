package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	mu sync.Mutex

	counter int
}

func (c *Counter) Count() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.counter++
}

func (c *Counter) getValue() int { // сделаем приватным метод (инкапсуляция, доступно только внутри пакета)
	return c.counter
}

func main() {
	c := &Counter{sync.Mutex{}, 0}

	wg := sync.WaitGroup{}
	wg.Add(10000)
	for i := 0; i < 10000; i++ {
		go func() {
			defer wg.Done()
			c.Count()
		}()
	}
	wg.Wait()

	fmt.Println(c.getValue())
}
