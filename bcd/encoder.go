package bcd

import (
	"fmt"
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
	var b byte
	for n = 0; n < len(p); n++ {
		b, err = Encode(p[n])
		if err != nil {
			return
		}
		n, err = enc.dst.Write([]byte{b})
		if err != nil {
			return
		}
	}

	return
}

// Encode TODO
func Encode(p byte) (byte, error) {
	if p > 99 {
		return 0xFF, fmt.Errorf("can't convert to BCD for p > 99")
	}

	b := (p / 10 << 4) + p%10

	return b, nil
}
