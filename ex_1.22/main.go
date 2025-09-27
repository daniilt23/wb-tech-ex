package main

import (
	"fmt"
	"math/big"
)

func main() {
	// создадим две переменные типа big.Int
	a := new(big.Int)
	b := new(big.Int)
	// зададим им значения превышающее 1 << 63
	a.SetString("3600000000000000000", 10)
	b.SetString("240000000000000000", 10)

	// используем функции библиотеки math/big для математических действий с числами любой длины
	fmt.Println("a + b = ", new(big.Int).Add(a, b))
	fmt.Println("a - b = ", new(big.Int).Sub(a, b))
	fmt.Println("a * b = ", new(big.Int).Mul(a, b))
	fmt.Println("a / b = ", new(big.Int).Div(a, b))
}
