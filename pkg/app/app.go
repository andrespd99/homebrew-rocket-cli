package app

import (
	"github.com/andrespd99/rocket-cli/pkg/generator"
)

type App struct {
	G generator.Generator
}

func NewApp() *App {
	return &App{
		G: generator.NewGenerator(),
	}
}
