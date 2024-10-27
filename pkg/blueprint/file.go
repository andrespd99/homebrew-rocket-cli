package blueprint

import (
	"encoding/base64"
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/andrespd99/rocket-cli/pkg/converters"
)

type Bundle []File

type File struct {
	Path          string `json:"path"`
	Base64Content string `json:"content"`
}

func BundleFromJSON(data []byte) (Bundle, error) {
	b := new(Bundle)

	if err := json.Unmarshal(data, b); err != nil {
		return nil, err
	}

	return *b, nil
}

// Create creates the file associated to this Data object with dst as the root directory
func (d *File) Create(params interface{}, dst string) error {
	content, err := base64.StdEncoding.DecodeString(d.Base64Content)
	if err != nil {
		return err
	}

	f, err := create(filepath.Join(dst, d.Path))
	if err != nil {
		return err
	}

	defer f.Close()

	if ignoreTemplate(f) {
		f.Write(content)

		return nil
	}

	if err = d.executeTemplate(content, params, f); err != nil {
		return err
	}

	return nil
}

func (d *File) executeTemplate(data []byte, params interface{}, file *os.File) error {
	var convertersFuncMap = template.FuncMap{
		"camelCase":  converters.ToCamelCase,
		"snakeCase":  converters.ToSnakeCase,
		"titleCase":  converters.ToTitleCase,
		"pascalCase": converters.ToPascalCase,
	}

	tmpl, err := template.New("").Funcs(convertersFuncMap).Parse(string(data))
	if err != nil {
		return err
	}

	err = tmpl.Execute(file, params)
	if err != nil {
		return err
	}

	return nil
}

func create(p string) (*os.File, error) {
	if err := os.MkdirAll(filepath.Dir(p), 0755); err != nil {
		return nil, err
	}
	return os.Create(p)
}

func ignoreTemplate(file *os.File) bool {
	for _, ext := range ignoredExtensions {
		if strings.HasSuffix(file.Name(), ext) {
			return true
		}
	}
	return false
}
