package main

import (
	"flag"
	"fmt"
	"path/filepath"
	"strings"
)

type FileExt string

const (
	NoExt  FileExt = ""
	AsmExt FileExt = ".s"
	HexExt FileExt = ".hex"
	TxtExt FileExt = ".txt"
)

type Flags struct {
	FlagSet flag.FlagSet
	Input   string
	Output  string
	Verbose bool
}

func NewFlags(name string) *Flags {

	fs := flag.NewFlagSet(name, flag.ContinueOnError)

	flags := &Flags{
		FlagSet: *fs,
	}

	fs.StringVar(&flags.Output, "output", "", "output filename")
	fs.StringVar(&flags.Output, "o", "", "output filename (shorthand)")

	fs.BoolVar(&flags.Verbose, "verbose", false, "verbosity of the assembler")
	fs.BoolVar(&flags.Verbose, "v", false, "verbosity of the assembler (shorthand)")

	return flags
}

func (f *Flags) Parse(rawArgs []string) error {
	// Parse flags
	err := f.FlagSet.Parse(rawArgs)

	if err != nil {
		return err
	}

	// Parse arguents
	args := f.FlagSet.Args()

	if len(args) != 1 {
		return fmt.Errorf("unexpected arguments %+v", args)
	}

	// Test inputfile
	input := args[0]
	ext := filepath.Ext(input)

	switch ext {
	case "", ".s":
		f.Input = input

	default:
		return fmt.Errorf("unknown file extension %v", ext)
	}

	return nil
}

func (f *Flags) SetOutput(ext FileExt) error {
	output := f.Output

	if output == "" {
		output = filepath.Base(f.Input)
	}

	f.Output = strings.Split(output, ".")[0] + string(ext)
	return nil
}
