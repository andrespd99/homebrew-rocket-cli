package main

import (
	"os"

	"github.com/andrespd99/rocket/pkg/blueprint/blueprints"
	"github.com/andrespd99/rocket/pkg/generator"
	"github.com/integrii/flaggy"
)

// Keep subcommands as globals so you can easily check if they were used later on.
var createSubcmd *flaggy.Subcommand
var flutterAppSubcmd *flaggy.Subcommand

var nameVar string
var appIdFlag string
var displayNameFlag string
var orgNameFlag = blueprints.DefaultOrgName
var descriptionFlag = blueprints.DefaultDescription
var platformsFlag = blueprints.DefaultPlatforms

func init() {
	flaggy.SetName("rocket")
	// TODO: Improve this mediocre description lmao
	flaggy.SetDescription("🚀 Rocket accelerates your Flutter project with out-of-the-box and production-ready code structure and configurations")

	// Create any subcommands and set their parameters.
	createSubcmd = flaggy.NewSubcommand("create")
	flutterAppSubcmd = flaggy.NewSubcommand("flutter_app")

	// TODO: Improve
	createSubcmd.Description = "Creates a project directory based on a pre-defined template"
	flutterAppSubcmd.Description = "Creates a Flutter project"

	flutterAppSubcmd.AddPositionalValue(&nameVar, "name", 1, true, "Package name")
	flutterAppSubcmd.String(&appIdFlag, "id", "app-id", "Reverse DNS identifier. Usually the Bundle ID (iOS) or Application ID (Android)")
	flutterAppSubcmd.String(&orgNameFlag, "org", "org-name", "Package name")
	flutterAppSubcmd.String(&displayNameFlag, "a", "disp-name", "Display name or label displayed on devices")
	flutterAppSubcmd.String(&descriptionFlag, "d", "desc", "Used in pubspec.yaml description and README overview")
	flutterAppSubcmd.StringSlice(&platformsFlag, "p", "platforms", "Target platforms")

	createSubcmd.AttachSubcommand(flutterAppSubcmd, 1)
	flaggy.AttachSubcommand(createSubcmd, 1)

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
				Name:        nameVar,
				AppId:       appIdFlag,
				OrgName:     orgNameFlag,
				DisplayName: displayNameFlag,
				Description: descriptionFlag,
				Platforms:   platformsFlag,
			})
			if err != nil {
				return err
			}

			g := generator.NewGenerator()

			g.GenerateAt(bp, nameVar)

			return nil
		}

		flaggy.ShowHelp("")

		return nil
	}
	flaggy.ShowHelp("")
	return nil
}
