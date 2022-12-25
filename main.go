package main

import (
	"flag"
	"fmt"
	"hthcrwzy/anySemicolon/interpreter"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	// Start REPL.
	// repl.Repl()

	flag.Parse()
	args := flag.Args()
	filename := args[0]

	code, err := ioutil.ReadFile(filename)

	if err != nil {
		log.Fatalf("Error: %v\n", err)
	}

	m := interpreter.NewMachine(string(code), os.Stdin, os.Stdout)
	m.Execute()

	fmt.Println() // 改行
}
