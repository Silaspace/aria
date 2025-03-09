package handler

import (
	"bufio"
	"bytes"
	"os"
)

const BUFFER_LEN int = 30

func NewFileReader(filename string) (*FileReader, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	r := bufio.NewReader(f)

	return &FileReader{
		File:   f,
		Reader: r,
	}, nil
}

func NewFileWriter(filename string) (*FileWriter, error) {
	f, err := os.Create(filename)
	if err != nil {
		return nil, err
	}

	w := bufio.NewWriter(f)

	return &FileWriter{
		File:   f,
		Writer: w,
	}, nil
}

func NewWebReader() *WebReader {
	r := bytes.NewReader([]byte{})
	return &WebReader{
		Reader: r,
	}
}

func NewWebWriter() *WebWriter {
	b := bytes.NewBuffer([]byte{})
	return &WebWriter{
		Buffer: b,
	}
}
