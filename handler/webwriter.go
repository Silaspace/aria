package handler

import (
	"bytes"
)

type WebWriter struct {
	Buffer *bytes.Buffer
}

func (r *WebWriter) Read() []byte {
	return r.Buffer.Bytes()
}

func (r *WebWriter) Write(data []byte) error {
	_, err := r.Buffer.Write(data)

	if err != nil {
		return err
	}

	return nil
}

func (r *WebWriter) Reset() {
	r.Buffer.Reset()
}

func (r *WebWriter) Close() {

}
