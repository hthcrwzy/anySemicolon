package main

import (
	"hthcrwzy/anySemicolon/interpreter"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	// Start REPL.
	// repl.Repl()

	filename := os.Args[1]
	code, err := ioutil.ReadFile(filename)

	if err != nil {
		log.Fatalf("Error: %v\n", err)
	}

	m := interpreter.NewMachine(string(code), os.Stdin, os.Stdout)
	m.Execute()
}
