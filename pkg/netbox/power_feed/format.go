package power_feed

import (
	"github.com/funtimecoding/soil/pkg/console/status"
	"github.com/funtimecoding/soil/pkg/console/status/option"
	"github.com/funtimecoding/soil/pkg/console/status/tag"
)

func (f *Feed) Format(o *option.Format) string {
	s := status.New(o)

	if o.HasTag(tag.Identifier) {
		s.Integer32(f.Identifier)
	}

	s.String(f.formatName(o)).RawList(f.Raw)

	return s.Format()
}
