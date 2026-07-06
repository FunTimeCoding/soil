package issue

import (
	"github.com/funtimecoding/soil/pkg/console/status"
	"github.com/funtimecoding/soil/pkg/console/status/option"
)

func (i *Issue) Format(f *option.Format) string {
	s := status.New(f).String(i.FormatStatus(), i.FormatSummary(f))

	if d := i.FormatDescription(f); d != NoDescription {
		s.Line("%s", d)
	}

	s.DetailLink(i.Link, i.Key, "")
	s.RawList(i)

	return s.Format()
}
