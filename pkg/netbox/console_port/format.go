package console_port

import (
	"github.com/funtimecoding/soil/pkg/console/constant"
	"github.com/funtimecoding/soil/pkg/console/status"
	"github.com/funtimecoding/soil/pkg/console/status/option"
)

func (p *Port) Format(f *option.Format) string {
	s := status.New(f)

	if f.HasTag(constant.TagIdentifier) {
		s.Integer32(p.Identifier)
	}

	s.String(p.formatName(f)).RawList(p.Raw)

	return s.Format()
}
