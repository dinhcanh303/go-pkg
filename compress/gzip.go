package compress

import (
	"bytes"
	"compress/gzip"
	"io"
)

const unzipLimit = 100 * 1024 * 1024 // 100MB

// Gzip compresses the input byte slice using gzip algorithm.
func Gzip(bs []byte) ([]byte, error) {
	var buf bytes.Buffer
	w := gzip.NewWriter(&buf)
	if _, err := w.Write(bs); err != nil {
		return nil, err
	}
	defer w.Close()

	return buf.Bytes(), nil
}

// Gunzip decompresses the input gzip-compressed byte slice.
func Gunzip(bs []byte) ([]byte, error) {
	r, err := gzip.NewReader(bytes.NewBuffer(bs))
	if err != nil {
		return nil, err
	}
	defer r.Close()

	var c bytes.Buffer
	if _, err = io.Copy(&c, io.LimitReader(r, unzipLimit)); err != nil {
		return nil, err
	}
	return c.Bytes(), nil
}
