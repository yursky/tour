package main

import "fmt"

var c, python, java bool

func main() {
	var i int
	fmt.Println(i, c, python, java)

	// var c, python, java = true, false, "no!"
	// fmt.Println(c, python, java)

	// c, python, java := true, false, "no!"
	// fmt.Println(c, python, java)
}

/*
*GOLANGS DEFAULT TYPES*

bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // alias for uint8

rune // alias for int32
     // represents a Unicode code point

float32 float64

complex64 complex128

*/
