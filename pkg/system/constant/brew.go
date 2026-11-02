package constant

import consoleConstant "github.com/funtimecoding/soil/pkg/console/constant"

const (
	Brew            = "brew"
	BrewOutdated    = "outdated"
	BrewInformation = "info"
	BrewInstalled   = "--installed"
	BrewNotation2   = "--json=v2"
	BrewNotation1   = "--json=v1"
)

var BrewFormat = consoleConstant.ColorFormat.Copy()
