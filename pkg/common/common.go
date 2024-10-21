package common

import (
	"io/fs"
)

// Commonly used things wrapped into one struct for convenience when passing it around
type Common struct {
	Debug bool
	Fs    fs.FS
}
