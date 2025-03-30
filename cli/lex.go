package main

import (
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
		exit(err)
	}

	err = flags.SetOutput(TxtExt)

	if err != nil {
		exit(err)
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
		exit(err)
	}

	writer, err := handler.NewFileWriter(lc.output)

	if err != nil {
		exit(err)
	}

	lex := lexer.NewLexer(reader)

	for {
		nextToken := lex.Next()
		writer.Write([]byte(nextToken.Fmt()))

		if nextToken.IsEOF() {
			break
		}

		if nextToken.IsErr() {
			break
		}
	}

	writer.Close()
	reader.Close()
}
