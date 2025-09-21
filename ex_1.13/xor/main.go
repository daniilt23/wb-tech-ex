package main

import "fmt"

func main() {
	a, b := 34, 56
	a = a ^ b
	b = a ^ b
	a = a ^ b
	fmt.Printf("%d %d\n", a, b)
}
