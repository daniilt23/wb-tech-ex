package outsideservice

import "fmt"

// создаем класс и реализацию какого то стороннего сервиса или библиотеки
// к которой у нас нет доступа (это может быть какой то сервис который что то оправляет на наш сервер)

type Target interface {
	Operation()
}

type ConcreteTagret struct {
}

func (t ConcreteTagret) Operation() {
	fmt.Println("Some move from new service")
}
