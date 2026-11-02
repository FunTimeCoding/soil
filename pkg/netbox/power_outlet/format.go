package power_outlet

import (
	consoleConstant "github.com/funtimecoding/soil/pkg/console/constant"
	"github.com/funtimecoding/soil/pkg/console/status"
	"github.com/funtimecoding/soil/pkg/console/status/option"
)

func (o *Outlet) Format(f *option.Format) string {
	s := status.New(f)

	if f.HasTag(consoleConstant.TagIdentifier) {
		s.Integer32(o.Identifier)
	}

	s.String(o.formatName(f)).RawList(o.Raw)

	return s.Format()
}
