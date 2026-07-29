package constant

import "github.com/funtimecoding/soil/pkg/console/status/option"

const (
	Brew            = "brew"
	BrewOutdated    = "outdated"
	BrewInformation = "info"
	BrewInstalled   = "--installed"
	BrewNotation2   = "--json=v2"
	BrewNotation1   = "--json=v1"
)

var BrewFormat = option.Color.Copy()
