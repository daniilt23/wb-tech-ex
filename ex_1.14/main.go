package main

import "fmt"

//не использую reflect так как нет необходимости в присваивании типа

func findType(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("value %d has type %T\n", v, v)
	case string:
		fmt.Printf("value %q has type %T\n", v, v)
	case bool:
		fmt.Printf("value %t has type %T\n", v, v)
	case chan struct{}: // сделаем для канала с пустой структурой для чистоты кода, если нужно все каналы то их можно перечислить
		fmt.Printf("channel with type %T\n", v)
	}
}

func main() {
	findType(21)
	findType("Hello")
	findType(true)
	ch := make(chan struct{}) 
	findType(ch)
}
