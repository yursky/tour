package main

import "fmt"

type AbstractInterface interface {
	Method()
}

type ConcreteType struct {
	String string
}

// This method means type ConcreteType implements the interface AbstractInterface,
// but we don't need to explicitly declare that it does so.
func (t ConcreteType) Method() {
	fmt.Println(t.String)
}

func main() {
	var example AbstractInterface = ConcreteType{"hello, world!"}
	example.Method()

	var i interface{}
	describe(i)

	i = 42
	describe(i)

	i = "hello"
	describe(i)
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
