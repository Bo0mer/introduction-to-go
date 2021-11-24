package main

import (
	"fmt"
	"time"
)

func main() {
	// START OMIT
	var zeroes = make(chan int)
	var ones = make(chan int)

	go func() {
		for {
			ones <- 1
		}
	}()
	go func() {
		for {
			zeroes <- 0
		}
	}()

	for {
		select {
		case x := <-ones:
			fmt.Print(x)
		case y := <-zeroes:
			fmt.Print(y)
		}
		time.Sleep(200 * time.Millisecond)
	}
	// END OMIT
}
