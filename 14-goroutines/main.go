package main

import (
	"fmt"
	// "sync"
	"time"
)

// var wg sync.WaitGroup

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
	// wg.Done()
}

func main() {
	// wg.Add(1)
	// go say("world") - could be a blocking API call
	say("hello")
	// wg,Wait()
}
