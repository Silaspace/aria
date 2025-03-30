package main

import (
	"os"

	"github.com/silaspace/aria/handler"
	"github.com/silaspace/aria/lexer"
	"github.com/silaspace/aria/parser"
)

type ParseCommand struct {
	input   string
	output  string
	verbose bool
}

func NewParseCommand(rawArgs []string) *ParseCommand {
	// Parse command line arguments
	flags := NewFlags("parse")
	err := flags.Parse(rawArgs)

	if err != nil {
		panic(err)
	}

	err = flags.SetOutput(TxtExt)

	if err != nil {
		panic(err)
	}

	// Return command
	return &ParseCommand{
		input:   flags.Input,
		output:  flags.Output,
		verbose: flags.Verbose,
	}
}

func (pc *ParseCommand) Run() {
	reader, err := handler.NewFileReader(pc.input)

	if err != nil {
		panic(err)
	}

	writer, err := handler.NewFileWriter(pc.output)

	if err != nil {
		panic(err)
	}

	lex := lexer.NewLexer(reader)
	parse := parser.NewParser(lex)

	for {
		nextLine := parse.Next()
		writer.Write([]byte(nextLine.Fmt()))

		if nextLine.Type() == parser.EOFType {
			break
		}

		if nextLine.Type() == parser.ErrorType {
			break
		}
	}

	writer.Close()
	reader.Close()
	os.Exit(0)
}
