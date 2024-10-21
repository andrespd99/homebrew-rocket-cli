package main

import (
	"log"

	"github.com/andrespd99/rocket-cli/pkg/generator"
	tmpl "github.com/andrespd99/rocket-cli/pkg/templates"
)

type ProjectData struct {
	ProjectName string
	Author      string
}

func main() {
	g := generator.NewGenerator()
	fileReader, err := tmpl.TestTemplate.Open()
	if err != nil {
		log.Fatalln(err)
	}
	defer fileReader.Close()

	// TODO: ADD PATH FLAGGY !

	dst := "test/"

	err = g.GenerateAt(fileReader, dst)
	if err != nil {
		log.Fatalln(err)
	}
}
