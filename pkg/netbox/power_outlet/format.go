package power_outlet

import (
	"github.com/funtimecoding/soil/pkg/console/status"
	"github.com/funtimecoding/soil/pkg/console/status/option"
	"github.com/funtimecoding/soil/pkg/console/status/tag"
)

func (o *Outlet) Format(f *option.Format) string {
	s := status.New(f)

	if f.HasTag(tag.Identifier) {
		s.Integer32(o.Identifier)
	}

	s.String(o.formatName(f)).RawList(o.Raw)

	return s.Format()
}
