package generator

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io"
	"os"
	"path"
	"strings"

	"github.com/andrespd99/rocket-cli/pkg/blueprints"
	"github.com/andrespd99/rocket-cli/pkg/types"
)

var ignoredExtensions = []string{".jpg", ".png", ".svg", ".gif", ".jpeg", ".webp", ".ico", ".icns"}

type Generator interface {
	Generate(blueprints []blueprints.Blueprint) error
	GenerateAt(blueprints []blueprints.Blueprint, dst string) error
}

type generator struct {
}

func NewGenerator() *generator {
	return &generator{}
}

// Generate generates the given tmpl files in the root directory.
func (g *generator) Generate(blueprints []blueprints.Blueprint) error {
	return g.GenerateAt(blueprints, "./")
}

// GenerateAt generates the given tmpl files at dst
func (g *generator) GenerateAt(blueprints []blueprints.Blueprint, dst string) error {
	for _, b := range blueprints {
		err := g.generate(b, dst)
		if err != nil {
			return err
		}
	}
	return nil
}

func (g *generator) generate(bprint blueprints.Blueprint, root string) error {
	buf := bytes.Buffer{}

	r, err := bprint.Open()
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

		if !shouldApplyTemplate(file) {
			continue
		}

		err = bprint.Execute(string(data), file)
		if err != nil {
			// os.Remove(dst)
			return fmt.Errorf("error when creating file %%: %s", b.Path, err.Error())
		}
	}

	return nil
}

func shouldApplyTemplate(file *os.File) bool {
	for _, ext := range ignoredExtensions {
		if strings.HasSuffix(file.Name(), ext) {
			return false
		}
	}
	return true
}
