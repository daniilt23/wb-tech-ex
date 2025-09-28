package main

import (
	"fmt"
	"strings"
)

func checkUnique(s string) bool {
	set := make(map[rune]struct{})
	s = strings.ToLower(s)
	for _, val := range s {
		_, ok := set[val]
		if ok {
			return false
		}
		set[val] = struct{}{}
	}

	return true
}

func main() {
	s1 := "qwertyQ"
	s2 := "abcdef"
	fmt.Println(checkUnique(s1)) // false
	fmt.Println(checkUnique(s2)) // true
}
