package lint

import "github.com/funtimecoding/soil/pkg/markup/scalar_or_list"

type Declaration struct {
	Base     scalar_or_list.Strings `yaml:"base"`
	Hosts    scalar_or_list.Strings `yaml:"hosts"`
	Commands scalar_or_list.Strings `yaml:"commands"`
}
