package report

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/text/report/line"
)

func (s *Section) Integer(
	name string,
	i int,
) {
	l := line.New()
	l.Value = fmt.Sprintf("%s: %d", name, i)
	l.Indent = s.indent + 1
	s.appendRenderable(l)
}
