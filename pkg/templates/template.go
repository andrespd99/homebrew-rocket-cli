package templates

import (
	"io"
	"os"
)

const tmplPath = "./templates"

type Template struct {
	Path string
	Data any
}

func (s Template) Open() (io.ReadCloser, error) {
	f, err := os.Open(s.Path)
	if err != nil {
		return nil, err
	}
	return f, nil
}
