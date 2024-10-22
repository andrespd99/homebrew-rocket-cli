package templates

import "fmt"

var TestTemplate = Template{
	Path: fmt.Sprintf("%s/base.json", tmplPath),
	Data: struct {
		Name        string
		Description string
	}{Name: "Avila Tek", Description: "Avila Tek Example description"},
}
