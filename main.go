package main

import (
	"log"
)

type ProjectData struct {
	ProjectName string
	Author      string
}

func main() {

	// TODO: ADD PATH FLAGGY !

	// dst := ".test/"

	if err := ServeCommand(); err != nil {
		log.Fatalln(err)
	}

	// bp, err := blueprints.NewFlutterAppBlueprint(blueprints.BaseFlutterAppParams{
	// 	Name:    "avila_tek_app",
	// 	OrgName: "com.avilatek",
	// })
	// if err != nil {
	// }

	// if err = app.G.GenerateAt(bp, dst); err != nil {
	// 	log.Fatalln(err)
	// }
}
