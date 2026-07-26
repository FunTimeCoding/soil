package power_feed

import (
	"github.com/funtimecoding/soil/pkg/console/constant"
	"github.com/funtimecoding/soil/pkg/console/status"
	"github.com/funtimecoding/soil/pkg/console/status/option"
)

func (f *Feed) Format(o *option.Format) string {
	s := status.New(o)

	if o.HasTag(constant.TagIdentifier) {
		s.Integer32(f.Identifier)
	}

	s.String(f.formatName(o)).RawList(f.Raw)

	return s.Format()
}
