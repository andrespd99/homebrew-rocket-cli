package templates

import (
	"io"
	"os"
)

const tmplPath = "./templates"

type template struct {
	Path string
}

func (s template) Open() (io.ReadCloser, error) {
	f, err := os.Open(s.Path)
	if err != nil {
		return nil, err
	}
	return f, nil
}
