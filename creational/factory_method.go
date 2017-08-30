package main

import (
	"flag"
	"fmt"
	"os"
)

// product
type Hamburger interface {
	Deliver()
}

type KfcHamburger struct{}

func (h KfcHamburger) Deliver() {
	fmt.Println("This is a hamburger from KFC.\n")
}

type McdonaldsHamburger struct{}

func (h McdonaldsHamburger) Deliver() {
	fmt.Println("This is a hamburger from McDonalds.\n")
}

// factory
type HamburgerFactory interface {
	Create() Hamburger
}

type Kfc struct{}

func (f Kfc) Create() Hamburger {
	return new(KfcHamburger)
}

type Mcdonalds struct{}

func (f Mcdonalds) Create() Hamburger {
	return new(McdonaldsHamburger)
}

//client
func main() {
	prefer := getPreferHamburger()

	var factory HamburgerFactory
	switch prefer {
	case "KFC":
		factory = new(Kfc)
	case "McDonalds":
		factory = new(Mcdonalds)
	default:
		fmt.Printf("%s not supported yet.\n", prefer)
		os.Exit(1)
	}

	hamburger := factory.Create()
	hamburger.Deliver()
}

func getPreferHamburger() string {

	var prefer string
	flag.StringVar(&prefer, "prefer", "KFC", "Preferenced Hamburger")
	flag.Parse()

	if len(os.Args) < 2 {
		flag.Usage()
		os.Exit(1)
	}

	return prefer
}
