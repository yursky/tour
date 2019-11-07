package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	fmt.Println(add(42, 13))

	a, b := swap("hello", "world")
	fmt.Println(a)
	fmt.Println(b)

	// to bypass a return variable
	// a, _ = swap("hello", "world")
	// fmt.Println(a)

	//IIFE, anonymous, and closure
	//func() {
	//	fmt.Println("Anonymous function and Closure demo")
	//}()
}
