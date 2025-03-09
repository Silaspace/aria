package main

import (
	"syscall/js"

	"github.com/silaspace/aria/assembler"
	"github.com/silaspace/aria/handler"
)

func main() {
	reader := handler.NewWebReader()
	writer := handler.NewWebWriter()
	asm := assembler.NewAssembler(reader, writer)

	js.Global().Set("write", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		data := make([]byte, args[0].Get("length").Int())
		js.CopyBytesToGo(data, args[0])
		reader.Write(data)
		return nil
	}))

	js.Global().Set("assemble", js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		writer.Reset()
		err := asm.Run()

		if err != nil {
			return js.Global().Get("Error").New(err.Error())
		}

		data := writer.Read()

		jsArray := js.Global().Get("Uint8Array").New(len(data))
		for i, b := range data {
			jsArray.SetIndex(i, b)
		}

		return jsArray
	}))

	select {}
}
