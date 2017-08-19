package main

import (
	"github.com/pirlo-san/design-patterns-in-go/oo"
)

func main() {
	p := oo.Person{}
	p.Init("pirlo", 21)
	p.SayHi()

	e := oo.Employee{}
	e.Init("kaka", 22, "milan")
	e.SayHi()
	e.Work()
}
