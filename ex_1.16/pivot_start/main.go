package main

// здесь я реализовал как вы и просили по условию чтобы опорный элемент был вначале но это очень неоптимально

import "fmt"

func quickSort(arr []int, low, high int) []int {
	if low < high {
		pivot := partition(arr[:], low, high)
		quickSort(arr, pivot+1, high)
		quickSort(arr, low, pivot-1)
	}

	return arr
}

func partition(arr []int, low, high int) int {
	pivot := arr[low] // обозначим pivot вначале массива (я бы делал медиану трех но по условию нельзя)
	i := high + 1

	for j := high; j > low; j-- {
		if arr[j] > pivot {
			i--
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	arr[i-1], arr[low] = arr[low], arr[i-1]

	return i - 1
}

func main() {
	mas := [15]int{2, 6, 5, 8, 9, 1, 3, 4, 7, 10, 17, 20, 11, 12, 13}
	sorted_mas := quickSort(mas[:], 0, len(mas)-1)
	fmt.Println(sorted_mas)
}
