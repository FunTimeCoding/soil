package virtual_chassis

import (
	"github.com/funtimecoding/soil/pkg/console/constant"
	"github.com/funtimecoding/soil/pkg/console/status"
	"github.com/funtimecoding/soil/pkg/console/status/option"
)

func (c *Chassis) Format(f *option.Format) string {
	s := status.New(f)

	if f.HasTag(constant.TagIdentifier) {
		s.Integer32(c.Identifier)
	}

	s.String(c.formatName(f)).RawList(c.Raw)

	return s.Format()
}
