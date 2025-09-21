package main

import "fmt"

func main() {
	a, b := 56, 34
	a, b = a + (b - a), b - (b - a)
	fmt.Printf("%d %d\n", a, b)
}