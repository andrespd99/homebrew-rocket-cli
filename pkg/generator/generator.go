package generator

import (
	"bytes"
	"encoding/base64"
	"io"
	"os"
	"path"

	"github.com/andrespd99/rocket-cli/pkg/templates"
	"github.com/andrespd99/rocket-cli/pkg/types"
)

type Generator interface {
	Generate(tmpl templates.Template) error
	GenerateAt(tmpl templates.Template, dst string) error
}

type generator struct {
}

func NewGenerator() *generator {
	return &generator{}
}

// Generate generates the given tmpl files in the root directory.
func (g *generator) Generate(tmpl templates.Template) error {
	return g.generate(tmpl, "./")
}

// GenerateAt generates the given tmpl files at dst
func (g *generator) GenerateAt(tmpl templates.Template, dst string) error {
	return g.generate(tmpl, dst)
}

func (g *generator) generate(tmpl templates.Template, root string) error {
	buf := bytes.Buffer{}

	r, err := tmpl.Open()
	if err != nil {
		return err
	}
	defer r.Close()

	_, err = io.Copy(&buf, r)
	if err != nil {
		return err
	}

	d := buf.Bytes()

	bundle, err := types.NewBundleFromJSON(d)
	if err != nil {
		return err
	}

	for _, b := range bundle {
		data, err := base64.StdEncoding.DecodeString(b.Content)
		if err != nil {
			return err
		}
		dst := path.Clean(path.Join(root, b.Path))

		err = os.MkdirAll(path.Dir(dst), 0755)
		if err != nil && !os.IsExist(err) {
			return err
		}

		file, err := os.Create(dst)
		if err != nil {
			return err
		}

		err = tmpl.Execute(string(data), file)
		if err != nil {
			os.Remove(dst)
			return err
		}
	}

	return nil
}
