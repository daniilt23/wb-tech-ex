package main

import "fmt"

func reverseString(s string) string {
	rune := make([]rune, len(s))
	n := 0

	for _, val := range s { // присваиваем в слайс рун символы
		rune[n] = val
		n++
	}

	for i := 0; i < n/2; i++ { // разворачиваем руны
		rune[i], rune[n-1-i] = rune[n-1-i], rune[i]
	}

	return string(rune)
}

func main() {
	s1 := "hello 犬"
	s2 := "главрыба"
	s1Reverse := reverseString(s1)
	s2Reverse := reverseString(s2)
	fmt.Println(s1Reverse) // 犬 olleh
	fmt.Println(s2Reverse) // абырвалг
}
