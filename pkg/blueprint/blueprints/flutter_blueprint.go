package blueprints

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/andrespd99/rocket-cli/pkg/blueprint"
	"github.com/andrespd99/rocket-cli/pkg/converters"
)

const (
	ios     = "ios"
	android = "android"
	web     = "web"
	macos   = "macos"
	linux   = "linux"
	windows = "windows"
	fuchsia = "fuchsia"

	baseFlutterAppJsonPath = "base.json"

	defaultOrgName     = "com.example"
	defaultDescription = "\"A Flutter project accelerated by Rocket CLI 🚀\""
)

var blueprintPaths = map[string]string{
	ios:     "ios.json",
	android: "android.json",
	web:     "web.json",
	macos:   "macos.json",
	linux:   "linux.json",
	windows: "windows.json",
	fuchsia: "fuchsia.json",
}

var (
	platforms        = []string{ios, android, web, macos, linux, windows, fuchsia}
	defaultPlatforms = []string{ios, android}
)

type BaseFlutterAppParams struct {
	// Organization name used in AppId
	OrgName string `json:"org-name"`
	// Reverse DNS name. Usually the Bundle ID (iOS) or Application ID (Android)
	AppId string `json:"app-id"`
	// Flutter package name
	Name string `json:"name"`
	// Name displayed on device app menus
	DisplayName string `json:"app-name"`
	// Used on pubspec.yaml description field and README overview.
	Description string `json:"description"`
	// Target platforms for the Flutter app.
	Platforms []string `json:"platforms"`
}

func flutterAppBlueprints(data BaseFlutterAppParams) ([]blueprint.Blueprint, error) {
	// TODO: SET FLUTTER_APP_PATH AND FLUTTER_ROOT_PATH (FOR IOS)

	var blueprints []blueprint.Blueprint
	data.Name = converters.ToSnakeCase(data.Name)

	if data.OrgName == "" {
		data.OrgName = defaultOrgName
	}

	if data.AppId == "" {
		data.AppId = fmt.Sprintf("%s.%s", data.OrgName, data.Name)
	}

	if data.DisplayName == "" {
		data.DisplayName = converters.ToTitleCase(data.Name)
	}

	if data.Description == "" {
		data.Description = defaultDescription
	}

	if len(data.Platforms) == 0 {
		data.Platforms = defaultPlatforms
	}

	// TODO: DEFINE AND PASS
	// FLUTTER_ROOT := ""
	// FLUTTER_APPLICATION_PATH := ""

	// NOTE: 🚨 template execution depends on struct's field names. * Renaming these fields is highly discoraged *. 🚨
	bpParams := struct {
		OrgName     string
		AppId       string
		Name        string
		AppName     string
		Description string
		// Name used for compiled files
		BuildableName string
		// Name used for compiled .exe files on Windows.
		BuildableNameWindows string
	}{
		OrgName:              data.OrgName,
		AppId:                data.AppId,
		Name:                 data.Name,
		AppName:              data.DisplayName,
		Description:          data.Description,
		BuildableName:        converters.ToPascalCase(data.Name),
		BuildableNameWindows: data.Name,
	}

	baseBp := blueprint.NewBlueprint(
		filepath.Join(jsonsPath, baseFlutterAppJsonPath),
		bpParams,
	)

	blueprints = append(blueprints, baseBp)

	for _, platform := range data.Platforms {
		blueprint, err := platformBlueprint(platform, bpParams)
		if err != nil {
			return blueprints, err
		}
		blueprints = append(blueprints, blueprint)
	}

	return blueprints, nil
}

func platformBlueprint(platform string, params interface{}) (blueprint.Blueprint, error) {
	platformFilePath, ok := blueprintPaths[platform]
	if !ok {
		platforms := strings.Join(platforms, ", ")
		err := fmt.Errorf("platform not supported: %s; supported platforms: %s", platform, platforms)
		return nil, err
	}

	return blueprint.NewBlueprint(filepath.Join(jsonsPath, platformFilePath), params), nil
}
