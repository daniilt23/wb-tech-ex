package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type Counter struct {
	counter int64
}

func (c *Counter) Count() {
	atomic.AddInt64(&c.counter, 1)
	// атомик в случаях где простой счетчик работает в раз быстрее мьютекса,
	// так как атомики работают на очень низком уровне (на некоторых ОС их реализации даже не видна в файлах го)
}

func (c *Counter) getValue() int64 {
	return c.counter
}

func main() {
	c := Counter{0}

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
