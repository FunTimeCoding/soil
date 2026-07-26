package tagged

import (
	"github.com/funtimecoding/soil/pkg/console/constant"
	"github.com/funtimecoding/soil/pkg/console/status"
	"github.com/funtimecoding/soil/pkg/console/status/option"
)

func (t *Tagged) Format(f *option.Format) string {
	s := status.New(f).Integer(t.Identifier).String(t.Name)

	if f.ShowExtended {
		s.TagLine(constant.TagUsage, "  line1")
		s.TagLine(constant.TagUsage, "  line2")
	}

	s.RawList(t.Raw)

	return s.Format()
}
