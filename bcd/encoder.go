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
		b[i] = (p[i] / 10 << 4) + p[i]%10
	}

	n, err = enc.dst.Write(b)

	return
}
