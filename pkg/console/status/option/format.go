package option

import "github.com/funtimecoding/soil/pkg/console/status/option/pair"

type Format struct {
	UseColor bool
	// Use short names, possibly acronyms
	UseCompact bool
	// Extended Additional multiple lines of information
	ShowExtended bool
	// Tags Toggles for fine-grained control of what to display
	Tags []string
	// Raw Print the wrapped object with %+v in a new line
	ShowRaw bool
	// Filters To filter arbitrary key-value pairs
	Filters     []*pair.Pair
	Indentation int
}
