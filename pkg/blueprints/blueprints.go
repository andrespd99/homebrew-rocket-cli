package blueprints

const jsonsPath = "./pkg/blueprints/jsons"

func NewFlutterAppBlueprint(data BaseFlutterAppData) ([]Blueprint, error) {
	return flutterAppBlueprints(data)
}
