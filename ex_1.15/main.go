package main

import "strings"

var justString string

func someFunc() {
	v := createHugeString(1024)
	// justString = v[:100] при таком присвоении у нас v не очистится пока не очистится justString
	// justString содержит в себе указатель и ее capacity будет равен capacity большой строки 
	// нам нужно присваивать 100 байтов в новую переменную которая не держит указатель на старый слайс чтобы он очистился
	str := make([]byte, 100)
	copy(str, v[:100]) // копируем данные в новую переменную
}

func createHugeString(size int) string {
	return strings.Repeat("x", size) // создаем строку весом 1024 байта
}

func main() {
	someFunc()
}
