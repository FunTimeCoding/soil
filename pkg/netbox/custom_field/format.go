package custom_field

import (
	consoleConstant "github.com/funtimecoding/soil/pkg/console/constant"
	"github.com/funtimecoding/soil/pkg/console/status"
	"github.com/funtimecoding/soil/pkg/console/status/option"
)

func (f *Field) Format(o *option.Format) string {
	s := status.New(o)

	if o.HasTag(consoleConstant.TagIdentifier) {
		s.Integer32(f.Identifier)
	}

	s.String(f.formatName(o)).RawList(f.Raw)

	return s.Format()
}
