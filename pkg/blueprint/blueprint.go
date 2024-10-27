package blueprint

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

// These file extensions are blacklisted to avoid template.Execute to fail due to
// the encoding format of some of these files; but also because it makes no sense to
// apply template values on these file types.
var ignoredExtensions = []string{".jpg", ".png", ".svg", ".gif", ".jpeg", ".webp", ".ico", ".icns"}

type blueprint struct {
	Path   string
	Params interface{}
}

func NewBlueprint(path string, params interface{}) Blueprint {
	return blueprint{
		Path:   path,
		Params: params,
	}
}

type Blueprint interface {
	// Execute creates all the files defined in the blueprint and writes the content
	Execute(dst string) error
}

func (bp blueprint) Execute(dst string) error {
	buf := bytes.Buffer{}

	r, err := os.Open(bp.Path)
	if err != nil {
		return err
	}

	defer r.Close()

	_, err = io.Copy(&buf, r)
	if err != nil {
		return err
	}

	d := buf.Bytes()

	bundle, err := BundleFromJSON(d)

	for _, b := range bundle {
		err := b.Create(bp.Params, dst)
		if err != nil {
			return err
		}
	}

	if err != nil {
		// os.Remove(dst)
		return fmt.Errorf("error when creating file %s: %s", bp.Path, err.Error())
	}

	return nil
}
