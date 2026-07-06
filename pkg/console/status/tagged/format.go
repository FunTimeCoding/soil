package tagged

import (
	"github.com/funtimecoding/soil/pkg/console/status"
	"github.com/funtimecoding/soil/pkg/console/status/option"
	"github.com/funtimecoding/soil/pkg/console/status/tag"
)

func (t *Tagged) Format(f *option.Format) string {
	s := status.New(f).Integer(t.Identifier).String(t.Name)

	if f.ShowExtended {
		s.TagLine(tag.Usage, "  line1")
		s.TagLine(tag.Usage, "  line2")
	}

	s.RawList(t.Raw)

	return s.Format()
}
