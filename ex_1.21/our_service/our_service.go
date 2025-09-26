package ourservice

import (
	"fmt"
	"ex_1.21/outside_service"
)

type Old struct { // класс который нужно адаптировать
}

func (o Old) AdapteeOpeartion() {
	fmt.Println("I am an adaptee operation")
}

type OldAdapter struct {
	old *Old
}

func createOld(o Old) Target {

}

func (adapter OldAdapter) Operation() {
	adapter.old.AdapteeOpeartion()
}
