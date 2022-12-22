package main

import (
	"flag"
	"fmt"
)

func main() {
	flag.Parse()
	args := flag.Args()

	fmt.Println("This program outputs decimal character codes.")
	fmt.Println()

	// args -> arg
	for _, v := range args {
		// string -> rune
		for _, char := range v {
			// rune -> int
			fmt.Printf("%v: %v\n", string(char), get_char_code(char))
		}
	}
}

func get_char_code(char rune) int {
	return int(char)
}
