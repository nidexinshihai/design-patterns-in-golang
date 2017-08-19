package oo

import (
	"fmt"
)

type Animal interface {
	Move()
	Shout()
}

type Dog struct {
}

func (dog Dog) Move() {
	fmt.Println("A dog moves with its legs.")
}

func (dog Dog) Shout() {
	fmt.Println("wang wang wang.")
}

type Bird struct {
}

func (bird Bird) Move() {
	fmt.Println("A bird flys with its wings.")
}

func (bird Bird) Shout() {
	fmt.Println("A bird shouts.")
}
