package main

import "fmt"

// сделаем ввод индекса с консоли и провалидируем ошибки
func readValue(index *int, len int) error {
	fmt.Println("Please enter index to delete:")
	if _, err := fmt.Scanf("%d", index); err != nil {
		return err
	}
	if *index > len || *index < 0 {
		return fmt.Errorf("index %d is out of range", *index)
	}

	return nil
}

func deleteElemByIndex(slice []int, index int) []int {
	copy(slice[index:], slice[index+1:])

	return slice[:len(slice)-1]
}

func main() {
	var index int
	slice := []int{1, 2, 3, 4, 5, 6, 7}
	lenSlice := len(slice)
	if err := readValue(&index, lenSlice); err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Printf("Slice before delete: %v len = %d cap = %d\n", slice, len(slice), cap(slice))
		// Slice before delete: [1 2 3 4 5 6 7] len = 7 cap = 7
		slice = deleteElemByIndex(slice, index)
		fmt.Printf("Slice after delete: %v len = %d cap = %d\n", slice, len(slice), cap(slice))
		// Slice after delete: [1 2 4 5 6 7] len = 6 cap = 7
	}
}
