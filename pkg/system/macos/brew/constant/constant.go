package constant

import "github.com/funtimecoding/soil/pkg/console/status/option"

const (
	Brew        = "brew"
	Outdated    = "outdated"
	Information = "info"
	Installed   = "--installed"
	Notation2   = "--json=v2"
	Notation1   = "--json=v1"
)

var Format = option.Color.Copy()
