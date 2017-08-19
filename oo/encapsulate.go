package oo

import (
	"fmt"
)

type Student struct {
	name  string
	age   int
	Major string
}

func (s *Student) Init(name string, age int, major string) {
	s.name = name
	s.age = age
	s.Major = major
}

func (s Student) SayHi() {
	fmt.Printf("Hi, I am %s aged %d, and my major is %s.\n", s.name, s.age, s.Major)
}
