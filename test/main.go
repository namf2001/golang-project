package main

import (
	"flag"
	"fmt"
)

func main() {
	// Define flags
	name := flag.String("name", "World", "a name to say hello to")
	age := flag.Int("age", 0, "your age")
	verbose := flag.Bool("verbose", false, "enable verbose mode")

	// Parse flags
	flag.Parse()

	// Use the flag values
	if *verbose {
		fmt.Printf("Verbose mode enabled\n")
	}
	fmt.Printf("Hello, %s!\n", *name)
	if *age > 0 {
		fmt.Printf("You are %d years old.\n", *age)
	}
}
