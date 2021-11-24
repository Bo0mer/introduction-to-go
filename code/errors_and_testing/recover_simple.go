package main

import (
	"fmt"
)

func foo() {
	panic("haha")
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered: %#v of type %T\n", r, r)
		} else {
			fmt.Println("Didn't recover anything, huh?")
		}
	}()

	foo()
}
