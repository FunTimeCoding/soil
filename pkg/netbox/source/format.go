package source

import (
	"github.com/funtimecoding/soil/pkg/console/constant"
	"github.com/funtimecoding/soil/pkg/console/status"
	"github.com/funtimecoding/soil/pkg/console/status/option"
)

func (s *Source) Format(f *option.Format) string {
	o := status.New(f)

	if f.HasTag(constant.TagIdentifier) {
		o.Integer32(s.Identifier)
	}

	o.String(s.formatName(f)).RawList(s.Raw)

	return o.Format()
}
