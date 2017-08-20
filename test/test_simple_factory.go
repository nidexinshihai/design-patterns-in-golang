package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/pirlo-san/design-patterns-in-go/creational"
)

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

func main() {

	prefer := getPreferHamburger()

	factory := new(creational.HamburgerFactory)
	hamburger := factory.CreateHamburger(prefer)
	if hamburger == nil {
		fmt.Printf("%s not supported yet.\n", prefer)
		os.Exit(1)
	}

	hamburger.Deliver()
}
