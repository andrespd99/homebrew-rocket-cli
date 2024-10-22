package main

import (
	"log"

	"github.com/andrespd99/rocket-cli/pkg/app"
	tmpl "github.com/andrespd99/rocket-cli/pkg/templates"
)

type ProjectData struct {
	ProjectName string
	Author      string
}

func main() {
	app := app.NewApp()

	// TODO: ADD PATH FLAGGY !

	dst := "test/"

	err := app.G.GenerateAt(tmpl.TestTemplate, dst)
	if err != nil {
		log.Fatalln(err)
	}
}
