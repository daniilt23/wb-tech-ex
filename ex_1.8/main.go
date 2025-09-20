package main

import (
	"fmt"
	"strconv"
)

func transform(value int64, i int, bit int) (result int64) {
	if bit == 1 {
		result = value | 1 << (i - 1) // так как самый правый бит это первый по условию то (i - 1)
	} else {
		result = value &^ (1 << (i - 1))
	}

	return
}

func main() {
	var value int64
	var i int
	var bit int

	fmt.Println("Please enter value:")
	if _, err := fmt.Scanf("%d", &value); err != nil { // валидация входного числа
		fmt.Println(err)
		return
	}

	fmt.Println("Please enter number of bit to swap:")
	if _, err := fmt.Scanf("%d", &i); err != nil { // валидация бита который хотим поменять
		fmt.Println(err)
		return
	}

	stringBinary := strconv.FormatInt(value, 2) // перевод числа в двоичную систему для проверки размера
	if i > len(stringBinary) {                  // проверяем что бит реально поменять в данном числе (не значащие нули не берем так как про них ничего не сказано)
		fmt.Println("The i bit is too large to our number")
		return
	}

	fmt.Println("Enter bit to change to")
	if _, err := fmt.Scanf("%d", &bit); err != nil { // валидация бита который хотим поменять
		fmt.Println(err)
		return
	}

	result := transform(value, i, bit)

	fmt.Printf("%d\n", result)
}
