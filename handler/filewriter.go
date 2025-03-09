package handler

import (
	"bufio"
	"os"
)

type FileWriter struct {
	File   *os.File
	Writer *bufio.Writer
}

func (w *FileWriter) Write(data []byte) error {
	_, err := w.Writer.Write(data)

	if err != nil {
		return err
	}

	return nil
}

func (w *FileWriter) Close() {
	w.Writer.Flush()
	w.File.Close()
}
