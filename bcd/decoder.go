package bcd

import (
	"fmt"
	"io"
)

type decoder struct {
	src io.Reader
}

// NewDecoder TODO
func NewDecoder(r io.Reader) io.Reader {
	return &decoder{
		src: r,
	}
}

// Read TODO
func (dec *decoder) Read(p []byte) (n int, err error) {
	b := make([]byte, len(p))

	n, err = dec.src.Read(b)
	if err != nil {
		return
	}

	m := n
	for n = 0; n < m; n++ {
		p[n], err = Decode(b[n])
		if err != nil {
			return
		}
	}

	return
}

// Decode TODO
func Decode(b byte) (byte, error) {
	p := ((b >> 4) * 10) + (b & 0x0f)

	if p > 99 {
		return 0xFF, fmt.Errorf("invalid BCD byte: overflow")
	}

	return p, nil
}
