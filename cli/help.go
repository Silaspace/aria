package main

import (
	"fmt"
)

const helptext = `

usage: aria command [options] inputfile
	command:
		lex		Output the tokenised input in a text file
		parse	Output the parsed syntax trees in a text file
		build	Fully assemble the input into a hex file
		help	Print this help menu
	options:
		-o, --output	Set the output file manually
		-v, --verbose	Increase the verbosity of the terminal output

`

type HelpCommand struct{}

func NewHelpCommand(rawArgs []string) *HelpCommand {
	return &HelpCommand{}
}

func (hc *HelpCommand) Run() {
	fmt.Print(helptext)
}
