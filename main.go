package main

import (
	"flag"
	"fmt"
)

func main() {
	// variables declaration

	tps := flag.Int("tps", 100, "an int")
	window := flag.Int("window", 20, "an int")

	flag.Parse() // after declaring flags we need to call it

	final := *tps + *window

	fmt.Println(final)

}
