package main

import (
	"log"

	"github.com/andrespd99/rocket-cli/pkg/app"
	"github.com/andrespd99/rocket-cli/pkg/blueprints"
	tmpl "github.com/andrespd99/rocket-cli/pkg/blueprints"
)

type ProjectData struct {
	ProjectName string
	Author      string
}

func main() {
	app := app.NewApp()

	// TODO: ADD PATH FLAGGY !

	dst := ".test/"

	bp, err := blueprints.NewFlutterAppBlueprint(tmpl.BaseFlutterAppData{
		Name:    "avila_tek_app",
		OrgName: "com.avilatek",
	})
	if err != nil {
		log.Fatalln(err)
	}

	if err = app.G.GenerateAt(bp, dst); err != nil {
		log.Fatalln(err)
	}
}
