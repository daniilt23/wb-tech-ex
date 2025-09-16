package main

import "fmt"

func main() {
	var workerCount int
	if _, err := fmt.Scanf("%d", &workerCount); err != nil {
		fmt.Println(err)
		return
	}

	ch := make(chan bool)
	for i := 0; i < workerCount; i++ {
		go func(ch chan bool) {
			for val := range ch {
				fmt.Printf("Worker %d take data %t\n", i, val)
			}
		}(ch)
	}

	for {
		ch <- true
	}
}
