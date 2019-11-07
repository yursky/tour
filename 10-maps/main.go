package main

import "fmt"

type Location struct {
	Lat, Long float64
}

var m map[string]Location

func main() {
	m = make(map[string]Location)

	// Insert or update an element
	m["Workday"] = Location{
		37.698543, -121.926455,
	}
	m["Austin"] = Location{
		30.267153, -97.7430608,
	}
	fmt.Println(m)

	// Delete an element
	delete(m, "Austin")
	fmt.Println(m)

	// Check for a values existence
	v, ok := m["Workday"]
	fmt.Println("The value:", v, "Present?", ok)
}
