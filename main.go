package main

import (
	"flag"
	"fmt"
)

func main() {
	source := flag.String("src", "en", "Source language to use for translation.")
	dest := flag.String("dest", "es", "Destination language for translation.")

	flag.Parse()

	for _, value := range flag.Args() {
		fmt.Println(value)
	}

	fmt.Printf("%v %v", *source, *dest)
}
