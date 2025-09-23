package main

import "fmt"

// использование метода медианы из 3 поможет избежать O(n^2) тк мы мы гарантированно
// никогда не используем худший элемент раздела в качестве опорного

func findPivotIndex(mas []int, low, high int) int {
	middle := (low + high) / 2
	first, second, third := mas[low], mas[middle], mas[high]

	if (first > second) != (first > third) {
		return low
	} else if (second > first) != (second > third) {
		return middle
	} else {
		return high
	}
}

func quickSort(mas []int, low, high int) []int {
	if low < high {
		pivot := partition(mas, low, high)
		quickSort(mas, pivot+1, high)
		quickSort(mas, low, pivot-1)
	}

	return mas
}

func partition(mas []int, low, high int) int {
	pivotIndex := findPivotIndex(mas, low, high)
	pivotValue := mas[pivotIndex]

	mas[pivotIndex], mas[high] = mas[high], mas[pivotIndex] // уберем пивот в конец для облегчения и разнообразия

	i := low - 1
	for j := low; j < high; j++ {
		if mas[j] < pivotValue {
			i++
			mas[j], mas[i] = mas[i], mas[j]
		}
	}

	mas[i+1], mas[high] = mas[high], mas[i+1]
	return i + 1
}

func main() {
	mas := [15]int{2, 6, 5, 8, 9, 1, 3, 13, 7, 10, 17, 20, 11, 12, 4}
	sortedMas := quickSort(mas[:], 0, len(mas)-1)
	fmt.Println(sortedMas)
}
