package main

import "fmt"

func main() {
	original := []string{"cat", "cat", "dog", "cat", "tree"}

	set := make(map[string]struct{})
	result := []string{}

	for _, val := range original {
		if _, exists := set[val]; !exists {
			result = append(result, val)
		}
		set[val] = struct{}{}
	}

	fmt.Println(result)
}
