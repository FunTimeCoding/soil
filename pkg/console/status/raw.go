package status

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/console"
)

func (s *Status) Raw(
	a any,
	title string,
) *Status {
	if !s.format.ShowRaw {
		return s
	}

	r := fmt.Sprintf("%+v", a)

	if r == "<nil>" {
		return s
	}

	if s.format.UseColor {
		r = console.Magenta("%s", r)
	}

	return s.Line("  %s: %s", title, r)
}
