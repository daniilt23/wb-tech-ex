package main

import "fmt"

func main() {
	sl1 := []int{1, 2, 3, 4, 5, 6, 7}
	sl2 := []int{2, 3, 4, 7}
	set := make(map[int]struct{})
	final := []int{}          // создаем слайс для ответа
	for _, val := range sl1 { // добавляем первый слайс в мапу
		set[val] = struct{}{}
	}

	for _, val := range sl2 {
		if _, exists := set[val]; exists { // проверяем наличие ключей, если такие есть то повторение
			final = append(final, val)
		}
	}

	fmt.Println(final) // [2 3 4 7]
}
