package main

import "fmt"

type Human struct {
	Name string
	Age  int
}

func (h *Human) Greeting() {
	fmt.Printf("Hello my name is %s and im %d years old\n", h.Name, h.Age)
}

type Action struct {
	Human
	Move string
}

func (a *Action) DoAction() {
	fmt.Printf("Hello my name is %s and im %d years old and im %s\n", a.Name, a.Age, a.Move)
}

func main() {
	h := Human{"Daniil", 19}
	h.Greeting() //Hello my name is Daniil and im 19 years old
	a := Action{h, "jump"}
	a.DoAction() //Hello my name is Daniil and im 19 years old and im jump
}
