package handler

import (
	"bufio"
	"io"
	"os"
)

type FileReader struct {
	File   *os.File
	Reader *bufio.Reader
}

func (r *FileReader) Next() (rune, error) {
	nextRune, _, err := r.Reader.ReadRune()

	if err == io.EOF {
		return 0, io.EOF
	}

	if err != nil {
		return 0, err
	}

	return nextRune, nil
}

func (r *FileReader) Peek() (rune, error) {
	nextRune, _, err := r.Reader.ReadRune()

	if err == io.EOF {
		return 0, io.EOF
	}

	if err != nil {
		return 0, err
	}

	err = r.Reader.UnreadRune()

	if err != nil {
		return 0, err
	}

	return nextRune, nil
}

func (r *FileReader) Reset() error {
	_, err := r.File.Seek(0, 0)

	if err != nil {
		return err
	}

	r.Reader.Reset(r.File)
	return nil
}

func (r *FileReader) Close() {
	r.File.Close()
}
