package creational

import (
	"fmt"
)

type Hamburger interface {
	Deliver()
}

type KfcHamburger struct{}

func (h KfcHamburger) Deliver() {
	fmt.Println("This is a hamburger from KFC.")
}

type McdonaldsHamburger struct{}

func (h McdonaldsHamburger) Deliver() {
	fmt.Println("This is a hamburger from McDonalds.")
}

type HamburgerFactory struct{}

func (f HamburgerFactory) CreateHamburger(prefer string) Hamburger {
	switch prefer {
	case "KFC":
		return new(KfcHamburger)
	case "McDonalds":
		return new(McdonaldsHamburger)
	default:
		return nil
	}
}
