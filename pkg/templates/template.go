package templates

import (
	"io"
	"os"
	"text/template"
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

func (t Template) Execute(data string, dst *os.File) error {

	funcMap := template.FuncMap{
		"camelCase": toCamelCase,
		"snakeCase": toSnakeCase,
	}

	tmpl, err := template.New("").Funcs(funcMap).Parse(data)
	if err != nil {
		return err
	}

	err = tmpl.Execute(dst, t.Data)
	if err != nil {
		return err
	}

	return nil
}
