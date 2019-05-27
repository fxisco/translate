package main

import (
	"flag"
	"fmt"
)

func main() {
	source := flag.String("src", "en", "Source language to use for translation.")
	dest := flag.String("dest", "es", "Destination language for translation.")

	flag.Parse()

	fmt.Printf("%v %v", source, dest)
}
