package main

import (
	"flag"
	"os"
	"path/filepath"
	"strings"

	"github.com/silaspace/aria/assembler"
	"github.com/silaspace/aria/handler"
	"github.com/silaspace/aria/lexer"
	"github.com/silaspace/aria/parser"
)

type Flags struct {
	Output  string
	Verbose bool
}

var output string
var verbose bool

func build(args []string, flags *Flags) {
	reader, err := handler.NewFileReader(args[0])

	if err != nil {
		panic(err)
	}

	writer, err := handler.NewFileWriter(flags.Output)

	if err != nil {
		panic(err)
	}

	asm := assembler.NewAssembler(reader, writer)
	err = asm.Run()

	if err != nil {
		panic(err)
	}

	asm.Close()
	os.Exit(0)
}

func lex(args []string, flags *Flags) {
	reader, err := handler.NewFileReader(args[0])

	if err != nil {
		panic(err)
	}

	writer, err := handler.NewFileWriter(flags.Output)

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
			break
		}
	}

	writer.Close()
	reader.Close()
	os.Exit(0)
}

func parse(args []string, flags *Flags) {
	reader, err := handler.NewFileReader(args[0])

	if err != nil {
		panic(err)
	}

	writer, err := handler.NewFileWriter(flags.Output)

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

func parseOutputFile(output string, input string) string {
	if output == "" {
		output = filepath.Base(input)
	}

	switch filepath.Ext(output) {
	case "":
		return output + ".hex"
	case ".hex":
		return output
	case ".s":
		return strings.Split(output, ".")[0] + ".hex"
	default:
		panic("AAAAAAAAAAA")
	}
}

func main() {
	// Set up cli flags
	flag.StringVar(&output, "output", "", "output filename")
	flag.StringVar(&output, "o", "", "output filename (shorthand)")

	flag.BoolVar(&verbose, "verbose", false, "verbosity of the assembler")
	flag.BoolVar(&verbose, "v", false, "verbosity of the assembler (shorthand)")

	// Parse flags
	flag.Parse()

	// Run command
	build(flag.Args(), &Flags{
		Output:  parseOutputFile(output, flag.Args()[0]),
		Verbose: verbose,
	})
}
