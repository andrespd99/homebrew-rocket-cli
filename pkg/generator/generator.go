package generator

import "github.com/andrespd99/rocket-cli/pkg/blueprint"

type (
	generator struct{}
	Generator interface {
		// Generate generates the given tmpl files in the root directory.
		Generate(blueprints []blueprint.Blueprint) error
		// GenerateAt generates the given tmpl files at dst
		GenerateAt(blueprints []blueprint.Blueprint, dst string) error
	}
)

func NewGenerator() *generator {
	return &generator{}
}

func (g *generator) Generate(blueprints []blueprint.Blueprint) error {
	return g.GenerateAt(blueprints, "./")
}

func (g *generator) GenerateAt(blueprints []blueprint.Blueprint, dst string) error {
	for _, b := range blueprints {
		err := g.generate(b, dst)
		if err != nil {
			return err
		}
	}
	return nil
}

func (g *generator) generate(bp blueprint.Blueprint, dst string) error {
	return bp.Execute(dst)
}
