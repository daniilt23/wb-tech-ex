package main

import "fmt"

// в данном решении можно не меняя исходную строку просто а на основе аски таблицы вставлять
// все регистры английского алфавита (работаем только с ним потому что в условии не сказано)

func checkUnique(s string) bool {
	set := make(map[rune]struct{})
	for _, val := range s {
		_, ok := set[val]
		if ok {
			return false
		}
		set[val] = struct{}{}
		if val < 91 {
			set[val+32] = struct{}{}
		} else {
			set[val-32] = struct{}{}
		}
	}

	return true
}

func main() {
	s1 := "qwertyQ"
	s2 := "abcdef"
	fmt.Println(checkUnique(s1)) // false
	fmt.Println(checkUnique(s2)) // true
}
