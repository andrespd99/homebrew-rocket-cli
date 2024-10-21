package generator

import (
	"bytes"
	"encoding/base64"
	"io"
	"os"
	"path"

	"github.com/andrespd99/rocket-cli/pkg/types"
)

type generator struct{}

func NewGenerator() *generator {
	return &generator{}
}

func (g *generator) Generate(tmpl io.Reader) error {
	return g.generateAt(tmpl, "./")
}

func (g *generator) GenerateAt(tmpl io.Reader, dst string) error {
	return g.generateAt(tmpl, dst)
}

func (g *generator) generateAt(tmpl io.Reader, root string) error {
	buf := bytes.Buffer{}

	_, err := io.Copy(&buf, tmpl)
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

		dst := path.Clean(root + b.Path)

		if err := os.MkdirAll(path.Join(path.Dir(dst)), 0755); err != nil && !os.IsExist(err) {
			return err
		}

		err = os.WriteFile(dst, data, 0755)
		if err != nil {
			os.Remove(dst)
			return err
		}
	}

	return nil
}
