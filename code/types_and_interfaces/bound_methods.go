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
	// pGreeter is bound to p.
	pGreeter := p.SayHi

	p.SayHi("Pesho")
	pGreeter("Pesho")

	fmt.Printf("%T\n", pGreeter)
}
