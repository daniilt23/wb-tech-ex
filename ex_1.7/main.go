package main

import (
	"fmt"
	"sync"
	"time"
)

type SafeCounter struct {
	mu sync.Mutex
	m  map[string]int
}

func (sc *SafeCounter) safeIncrement(key string) {
	sc.mu.Lock() // блокируем для доступа только одной горутины
	sc.m[key]++
	sc.mu.Unlock() // разблокирование для прибавления счетчика
}

func (sc *SafeCounter) getValue(key string) int {
	return sc.m[key]
}

func main() {
	sc := SafeCounter{m: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go sc.safeIncrement("some_key")
	}

	time.Sleep(time.Second)              // сделаем ожидание чтобы горутины выполнились
	fmt.Println(sc.getValue("some_key")) // проверяем что все прибавилось правильно выведет 1000
}
