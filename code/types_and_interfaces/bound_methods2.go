package main

import (
	"fmt"
)

type greeter struct {
	Name string
}

func (g *greeter) SayHi(name string) {
	fmt.Printf("Hi, %s! I am %s (%p)\n", name, g.Name, g)
}

func main() {
	var p = &greeter{Name: "p"}

	unboundGreeter := (*greeter).SayHi
	fmt.Printf("%T\n", unboundGreeter)

	unboundGreeter(p, "Gosho")
	unboundGreeter(&greeter{Name: "other"}, "Maria")
}
