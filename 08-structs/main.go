package main

import "fmt"

type Thing struct {
	X int
	Y int
}

func main() {
	v := Thing{1, 2}
	v.X = 4
	fmt.Println(v.X)
}
