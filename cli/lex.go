package main

import (
	"fmt"

	"github.com/silaspace/aria/handler"
	"github.com/silaspace/aria/lexer"
)

type LexCommand struct {
	input   string
	output  string
	verbose bool
}

func NewLexCommand(rawArgs []string) *LexCommand {
	// Parse command line arguments
	flags := NewFlags("lex")
	err := flags.Parse(rawArgs)

	if err != nil {
		panic(err)
	}

	err = flags.SetOutput(TxtExt)

	if err != nil {
		panic(err)
	}

	// Return command
	return &LexCommand{
		input:   flags.Input,
		output:  flags.Output,
		verbose: flags.Verbose,
	}
}

func (lc *LexCommand) Run() {
	reader, err := handler.NewFileReader(lc.input)

	if err != nil {
		panic(err)
	}

	writer, err := handler.NewFileWriter(lc.output)

	if err != nil {
		panic(err)
	}

	lex := lexer.NewLexer(reader)

	for {
		nextToken := lex.Next()
		writer.Write([]byte(nextToken.Fmt()))

		if nextToken.IsEOF() {
			break
		}

		if nextToken.IsErr() {
			fmt.Printf("Unexpected error - %v", nextToken.Value)
			break
		}
	}

	writer.Close()
	reader.Close()
}
