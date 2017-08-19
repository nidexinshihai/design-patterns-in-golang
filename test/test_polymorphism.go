package main

import (
	"github.com/pirlo-san/design-patterns-in-go/oo"
)

func main() {
	var animal oo.Animal

	animal = oo.Dog{}
	animal.Move()
	animal.Shout()

	animal = oo.Bird{}
	animal.Move()
	animal.Shout()
}
