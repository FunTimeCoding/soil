package build

import (
	"github.com/funtimecoding/soil/pkg/build/option"
	"github.com/funtimecoding/soil/pkg/strings/split"
)

func Architectures(o *option.Build) {
	for _, t := range o.Targets {
		parts := split.Dash(t)
		o.OperatingSystem = parts[0]
		o.Architecture = parts[1]
		Go(o)
	}
}
