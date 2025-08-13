package chardet

import (
	"bytes"
	"io"
	"golang.org/x/net/html/charset"
)

func DetectAndDecode(r io.Reader) (*bytes.Reader, error) {
	decoded, err := charset.NewReader(r, "text/html")
	if err != nil {
		return nil, err
	}

	buf := new(bytes.Buffer)
	_, err = io.Copy(buf, decoded)
	if err != nil {
		return nil, err
	}

	return bytes.NewReader(buf.Bytes()), nil
}

