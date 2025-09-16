package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	arr := [5]int{2, 4, 6, 8, 10}

	//еще так ккк мы заранее знаем количество элементов массива
	//можно было написать заранее wg.Add(5)
	for _, val := range arr {
		wg.Add(1)
		go calculateSquare(val, &wg)
	}

	wg.Wait()
}

// название может быть и printSquare (так как делаем
// 2 действия, но декомпозировать на еще 2 функции в такой задаче не вижу смысла)
func calculateSquare(val int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(val * val)
}
