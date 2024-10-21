package app

import (
	"io"

	"github.com/andrespd99/rocket-cli/pkg/common"
)

type App struct {
	*common.Common
	closers []io.Closer
}
