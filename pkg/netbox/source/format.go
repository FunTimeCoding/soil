package source

import (
	consoleConstant "github.com/funtimecoding/soil/pkg/console/constant"
	"github.com/funtimecoding/soil/pkg/console/status"
	"github.com/funtimecoding/soil/pkg/console/status/option"
)

func (s *Source) Format(f *option.Format) string {
	o := status.New(f)

	if f.HasTag(consoleConstant.TagIdentifier) {
		o.Integer32(s.Identifier)
	}

	o.String(s.formatName(f)).RawList(s.Raw)

	return o.Format()
}
