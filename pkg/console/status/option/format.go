package option

import "github.com/funtimecoding/soil/pkg/console/status/option/pair"

type Format struct {
	UseColor     bool
	UseCompact   bool
	ShowExtended bool
	Tags         []string
	ShowRaw      bool
	Filters      []*pair.Pair
	Indentation  int
}
