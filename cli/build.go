package main

import (
	"github.com/silaspace/aria/assembler"
	"github.com/silaspace/aria/handler"
)

type BuildCommand struct {
	input   string
	output  string
	verbose bool
}

func NewBuildCommand(rawArgs []string) *BuildCommand {
	// Parse command line arguments
	flags := NewFlags("build")
	err := flags.Parse(rawArgs)

	if err != nil {
		exit(err)
	}

	err = flags.SetOutput(HexExt)

	if err != nil {
		exit(err)
	}

	// Return command
	return &BuildCommand{
		input:   flags.Input,
		output:  flags.Output,
		verbose: flags.Verbose,
	}
}

func (bc *BuildCommand) Run() {
	reader, err := handler.NewFileReader(bc.input)

	if err != nil {
		exit(err)
	}

	writer, err := handler.NewFileWriter(bc.output)

	if err != nil {
		exit(err)
	}

	asm := assembler.NewAssembler(reader, writer)
	err = asm.Run()

	if err != nil {
		exit(err)
	}

	asm.Close()
}
