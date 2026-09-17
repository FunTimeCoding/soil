package lint

import "github.com/funtimecoding/soil/pkg/markup/scalar_or_list"

type declaration struct {
	Base scalar_or_list.Strings `yaml:"base"`
}
