package bcd

import (
	"io"
)

type encoder struct {
	dst io.Writer
}

// NewEncoder TODO
func NewEncoder(w io.Writer) io.Writer {
	return &encoder{
		dst: w,
	}
}

// Write TODO
func (enc *encoder) Write(p []byte) (n int, err error) {
	b := make([]byte, len(p))

	for i := 0; i < len(p); i++ {
		b[i] = Encode(p[i])
	}

	n, err = enc.dst.Write(b)

	return
}

// Encode TODO
func Encode(p byte) byte {
	b := (p / 10 << 4) + p%10

	return b
}
