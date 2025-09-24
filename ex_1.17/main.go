package main

import "fmt"

func binarySearch(slice []int, value int) int {
	left, right := 0, len(slice)-1
	for left <= right {
		middle := (left + right) / 2

		if slice[middle] == value {
			return middle
		} else if slice[middle] < value {
			left = middle + 1
		} else {
			right = middle - 1
		}
	}

	return -1 // если ничего не нашли
}

func main() {
	slice := []int{2, 4, 5, 7, 9, 12, 23, 34, 56, 78, 99} // создадим сразу отсортированный
	// если нужно будет отсортировать то воспользуемся сортировкой из прошлого задания (1.16)
	index := binarySearch(slice, 34)
	undefinedIndex := binarySearch(slice, 6)
	fmt.Printf("%d %d\n", index, undefinedIndex) // 7 -1
}
