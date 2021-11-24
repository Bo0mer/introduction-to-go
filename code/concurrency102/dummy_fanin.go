package main

import (
	"fmt"
	"math/rand"
	"time"
)

func talk(msg string) <-chan string { // HL
	c := make(chan string)
	go func() {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s: %d", msg, i)
			time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
		}
	}()
	return c
}

func main() {
	bob := talk("Bob")     // HL
	alice := talk("Alice") // HL
	for i := 0; i < 5; i++ {
		fmt.Println(<-bob)
		fmt.Println(<-alice)
	}
}
