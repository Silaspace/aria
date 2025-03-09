package handler

import (
	"bytes"
	"io"
)

type WebReader struct {
	Reader *bytes.Reader
}

func (r *WebReader) Next() (rune, error) {
	nextRune, _, err := r.Reader.ReadRune()

	if err == io.EOF {
		return 0, io.EOF
	}

	if err != nil {
		return 0, err
	}

	return nextRune, nil
}

func (r *WebReader) Peek() (rune, error) {
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

func (r *WebReader) Write(data []byte) {
	r.Reader.Reset(data)
}

func (r *WebReader) Reset() error {
	_, err := r.Reader.Seek(0, 0)

	if err != nil {
		return err
	}

	return nil
}

func (r *WebReader) Close() {

}
