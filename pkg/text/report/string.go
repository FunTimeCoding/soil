package report

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/text/report/line"
)

func (s *Section) String(
	name string,
	v string,
) {
	l := line.New()
	l.Value = fmt.Sprintf("%s: %s", name, v)
	l.Indent = s.indent + 1
	s.appendRenderable(l)
}
