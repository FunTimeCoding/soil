package customer

import (
	"github.com/funtimecoding/soil/pkg/console/status"
	"github.com/funtimecoding/soil/pkg/console/status/option"
)

func (i *Issue) Format(f *option.Format) string {
	s := status.New(f).String(i.Key, i.Status, i.Title).RawList(i.Raw)

	if f.ShowRaw {
		s.Line("  RequestFieldValues: %+v", i.Raw.RequestFieldValues)
		s.Line("  Links: %+v", i.Raw.Links)
	}

	return s.Format()
}
