package main

import "fmt"

type Payment interface { // у нас есть интерфейс для оплаты
	ProcessPayment()
}

type PaymentProcessor struct{}

func (pp PaymentProcessor) ProcessPayment() {
	// любая бизнес логика
	fmt.Println("Make payment with old system")
}

// мы хотим интегрировать новый способ оплаты от внешнего сервиса который мы не можем менять
type NewPaymentProcessor struct{}

func (npp NewPaymentProcessor) NewProcessPayment() {
	// любая бизнес логика
	fmt.Println("Make payment with new system")
}

// создаем структуру которая включает полностью новый интерфейс
type NewPaymentAdapter struct {
	newPaymentProcessor *NewPaymentProcessor
}

func (npa NewPaymentAdapter) ProcessPayment() {
	npa.newPaymentProcessor.NewProcessPayment()
}

func main() {
	var p Payment // единый интерфейс для всех оплат

	p = &PaymentProcessor{}
	p.ProcessPayment() // Make payment with old system

	p = &NewPaymentAdapter{newPaymentProcessor: &NewPaymentProcessor{}}
	p.ProcessPayment() // Make payment with new system
}
