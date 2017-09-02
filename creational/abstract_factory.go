package main

import (
	"flag"
	"fmt"
	"os"
)

// product type one: hamburger
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

// product type two: chicken wing
type ChickenWing interface {
	Deliver()
}

type KfcChickenWing struct{}

func (wing KfcChickenWing) Deliver() {
	fmt.Println("This is a chicken wing from KFC.")
}

type McdonaldsChickenWing struct{}

func (wing McdonaldsChickenWing) Deliver() {
	fmt.Println("This is a chicken wing from McDonalds.")
}

// abstract factory
type FastFoodFactory interface {
	CreateHamburger() Hamburger
	CreateChickenWing() ChickenWing
}

// concrete factory one: KFC
type Kfc struct{}

func (factory Kfc) CreateHamburger() Hamburger {
	return new(KfcHamburger)
}
func (factory Kfc) CreateChickenWing() ChickenWing {
	return new(KfcChickenWing)
}

// concrete factory two: McDonalds
type Mcdonalds struct{}

func (factory Mcdonalds) CreateHamburger() Hamburger {
	return new(McdonaldsHamburger)
}
func (factory Mcdonalds) CreateChickenWing() ChickenWing {
	return new(McdonaldsChickenWing)
}

func main() {
	prefer := GetPreferredFastFood()

	var factory FastFoodFactory
	switch prefer {
	case "KFC":
		factory = new(Kfc)
	case "McDonalds":
		factory = new(Mcdonalds)
	default:
		fmt.Printf("%s not supported yet.\n", prefer)
		os.Exit(1)
	}

	hamburger := factory.CreateHamburger()
	chickenWing := factory.CreateChickenWing()

	hamburger.Deliver()
	chickenWing.Deliver()
}

func GetPreferredFastFood() string {

	var prefer string
	flag.StringVar(&prefer, "prefer", "KFC", "preferred fast food type")
	flag.Parse()

	if len(os.Args) != 3 {
		flag.Usage()
		os.Exit(1)
	}

	return prefer
}
