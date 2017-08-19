package oo

import (
	"fmt"
)

type Person struct {
	name string
	age  int
}

func (p *Person) Init(name string, age int) {
	p.name = name
	p.age = age
}

func (p Person) SayHi() {
	fmt.Printf("Hi, I am %s, %d years old.\n", p.name, p.age)
}

type Employee struct {
	Person
	company string
}

func (e *Employee) Init(name string, age int, company string) {
	e.Person.Init(name, age)
	e.company = company
}

func (e Employee) Work() {
	fmt.Printf("I'm working in %s.\n", e.company)
}
