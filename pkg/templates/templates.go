package templates

import "fmt"

var TestTemplate = template{
	Path: fmt.Sprintf("%s/base.json", tmplPath),
}
