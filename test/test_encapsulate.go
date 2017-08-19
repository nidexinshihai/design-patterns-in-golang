package main

import (
	"github.com/pirlo-san/design-patterns-in-go/oo"
)

func main() {
	s := oo.Student{}
	s.Init("pirlo", 21, "cs")
	s.SayHi()

	s.Major = "finance"
	s.SayHi()

	s.age = 22
	s.SayHi()
}
