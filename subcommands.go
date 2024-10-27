package main

import (
	"log"
	"os"

	"github.com/andrespd99/rocket-cli/pkg/blueprint/blueprints"
	"github.com/andrespd99/rocket-cli/pkg/generator"
	"github.com/integrii/flaggy"
)

// Keep subcommands as globals so you can easily check if they were used later on.
var createSubcmd *flaggy.Subcommand
var flutterAppSubcmd *flaggy.Subcommand

var nameVar string
var appIdFlag string
var orgNameFlag string
var displayNameFlag string
var descriptionFlag string
var platformsFlag string

func init() {
	flaggy.SetName("🚀 Rocket")
	// TODO: Improve this mediocre description lmao
	flaggy.SetDescription("Accelerate your Flutter project with out-of-the-box and production-ready setup")

	// Create any subcommands and set their parameters.
	createSubcmd = flaggy.NewSubcommand("create")
	flutterAppSubcmd = flaggy.NewSubcommand("flutter_app")

	createSubcmd.AddPositionalValue(&nameVar, "name", 1, true, "Package name")

	createSubcmd.AttachSubcommand(flutterAppSubcmd, 1)

	flaggy.AttachSubcommand(createSubcmd, 1)

	// TODO: Improve
	createSubcmd.Description = "Creates a project directory based on a pre-defined template"

	createSubcmd.String(&appIdFlag, "id", "app-id", "Reverse DNS identifier. Usually the Bundle ID (iOS) or Application ID (Android)")
	createSubcmd.String(&orgNameFlag, "org", "org-name", "Package name")
	createSubcmd.String(&displayNameFlag, "a", "disp-name", "Display name or label displayed on devices")
	createSubcmd.String(&descriptionFlag, "d", "desc", "Used in pubspec.yaml description and README overview")
	createSubcmd.String(&platformsFlag, "p", "platforms", "Target platforms")

	// Set the version and parse all inputs into variables.
	var version string

	b, err := os.ReadFile("./version")
	if err != nil {
		version = "unknown"
	} else {
		version = string(b)
	}

	flaggy.SetVersion(version)
	flaggy.Parse()
}

func ServeCommand() error {
	if createSubcmd.Used {
		if flutterAppSubcmd.Used {
			bp, err := blueprints.NewFlutterAppBlueprint(blueprints.BaseFlutterAppParams{
				Name: nameVar,
			})
			if err != nil {
				log.Fatalln(err)
			}

			g := generator.NewGenerator()
			g.Generate(bp)
		}
	}

	flaggy.ShowHelp("")

	return nil
}
