package assembler

import (
	"github.com/silaspace/aria/lexer"
	"github.com/silaspace/aria/parser"
)

func NewAssembler(reader Reader, writer Writer) *Assembler {
	l := lexer.NewLexer(reader)
	p := parser.NewParser(l)

	return &Assembler{
		Reader: reader,
		Writer: writer,
		Parser: *p,
		PC:     0,
	}
}
