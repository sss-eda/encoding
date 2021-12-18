package lemi025

import (
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

	for i := 0; i < n; i++ {
		p[i] = ((b[i] & 0xf0 >> 4) * 10) + (b[i] & 0x0f)
	}

	return
}
