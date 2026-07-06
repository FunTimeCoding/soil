package report

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/text/report/line"
)

func (s *Section) Float(
	name string,
	f float64,
	unit string,
) {
	var result string

	if unit == "" {
		result = fmt.Sprintf("%s: %.1f", name, f)
	} else {
		result = fmt.Sprintf("%s: %.1f %s", name, f, unit)
	}

	l := line.New()
	l.Value = result
	l.Indent = s.indent + 1
	s.appendRenderable(l)
}
