package assembler

import (
	"github.com/silaspace/aria/device"
	"github.com/silaspace/aria/lexer"
	"github.com/silaspace/aria/parser"
)

func NewAssembler(reader Reader, writer Writer) *Assembler {
	l := lexer.NewLexer(reader)
	p := parser.NewParser(l)
	d := device.DefaultDevice()

	return &Assembler{
		Device:  *d,
		Parser:  *p,
		PC:      0,
		Reader:  reader,
		Symbols: map[string]uint64{},
		Writer:  writer,
	}
}
