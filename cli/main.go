package main

import (
	"fmt"
	"os"
)

type Command interface {
	Run()
}

func exit(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	os.Exit(0)
}

func main() {
	subcommand := os.Args[1]

	switch subcommand {
	case "help":
		hc := NewHelpCommand(os.Args[2:])
		hc.Run()

	case "lex":
		lc := NewLexCommand(os.Args[2:])
		lc.Run()

	case "parse":
		pc := NewParseCommand(os.Args[2:])
		pc.Run()

	case "build":
		bc := NewBuildCommand(os.Args[2:])
		bc.Run()

	default:
		fmt.Printf("Invalid subcommand '%v' - try 'aria help' for more information\n", subcommand)
	}

	exit(nil)
}
