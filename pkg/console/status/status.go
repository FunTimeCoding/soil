package status

import (
	"github.com/funtimecoding/soil/pkg/console/status/option"
	"github.com/funtimecoding/soil/pkg/console/status/tag_line"
)

type Status struct {
	bubbles    []string
	lines      []string
	linesByTag []*tag_line.Line
	format     *option.Format
}
