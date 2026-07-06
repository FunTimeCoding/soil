package extended

import (
	"github.com/funtimecoding/soil/pkg/console/status"
	"github.com/funtimecoding/soil/pkg/console/status/option"
)

func (e *Extended) Format(f *option.Format) string {
	s := status.New(f).Integer(e.Identifier).String(e.Name, e.Description)

	if f.ShowExtended {
		s.Line("  line1")
		s.Line("  line2")
	}

	s.RawList(e.Raw)

	return s.Format()
}
