package blueprints

import (
	"fmt"
	"strings"
)

const (
	ios     = "ios"
	android = "android"
	web     = "web"
	macos   = "macos"
	linux   = "linux"
	windows = "windows"
	fuchsia = "fuchsia"
)

var platforms = []string{ios, android, web, macos, linux, windows, fuchsia}

type BaseFlutterAppData struct {
	OrgName              string   `json:"org-name"`
	AppId                string   `json:"app-id"`
	Name                 string   `json:"name"`     // avila_tek
	AppName              string   `json:"app-name"` // Avila Tek
	Description          string   `json:"description"`
	Platforms            []string `json:"platforms"`
	BuildableName        string   //AvilaTek
	BuildableNameWindows string   //avila_tek
}

func flutterAppBlueprints(data BaseFlutterAppData) ([]Blueprint, error) {
	// TODO: SET FLUTTER_APP_PATH AND FLUTTER_ROOT_PATH (FOR IOS)
	var blueprints []Blueprint
	data.Name = toSnakeCase(data.Name)

	if data.OrgName == "" {
		data.OrgName = "com.example"
	}

	if data.AppId == "" {
		data.AppId = fmt.Sprintf("%s.%s", data.OrgName, data.Name)
	}

	if data.AppName == "" {
		data.AppName = toTitleCase(data.Name)
	}

	if data.Description == "" {
		data.Description = "\"A Flutter project accelerated by Rocket CLI 🚀\""
	}

	if len(data.Platforms) == 0 {
		data.Platforms = []string{ios, android}
	}

	data.BuildableName = toPascalCase(data.Name)
	data.BuildableNameWindows = data.Name

	// TODO: DEFINE AND PASS
	// FLUTTER_ROOT := ""
	// FLUTTER_APPLICATION_PATH := ""

	blueprints = append(
		blueprints,
		Blueprint{
			Path: fmt.Sprintf("%s/base.json", jsonsPath),
			Data: data,
		})

	for _, platform := range data.Platforms {
		blueprint, err := platformBlueprint(platform, data)
		if err != nil {
			return blueprints, err
		}
		blueprints = append(blueprints, blueprint)
	}

	return blueprints, nil
}

func platformBlueprint(platform string, data BaseFlutterAppData) (Blueprint, error) {
	var (
		path      string
		blueprint Blueprint
	)

	switch platform {
	case ios:
		path = fmt.Sprintf("%s/ios.json", jsonsPath)
	case android:
		path = fmt.Sprintf("%s/android.json", jsonsPath)
	case web:
		path = fmt.Sprintf("%s/web.json", jsonsPath)
	case macos:
		path = fmt.Sprintf("%s/macos.json", jsonsPath)
	case linux:
		path = fmt.Sprintf("%s/linux.json", jsonsPath)
	case windows:
		path = fmt.Sprintf("%s/windows.json", jsonsPath)
	case fuchsia:
		path = fmt.Sprintf("%s/fuchsia.json", jsonsPath)
	default:
		platforms := strings.Join(platforms, ", ")
		err := fmt.Errorf("platform not supported: %s; supported platforms: %s", platform, platforms)
		return blueprint, err
	}

	blueprint = Blueprint{
		Path: path,
		Data: data,
	}

	return blueprint, nil
}
