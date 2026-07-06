package prefix

import (
	"github.com/funtimecoding/soil/pkg/console/status"
	"github.com/funtimecoding/soil/pkg/console/status/option"
	"github.com/funtimecoding/soil/pkg/console/status/tag"
)

func (p *Prefix) Format(f *option.Format) string {
	s := status.New(f)

	if f.HasTag(tag.Identifier) {
		s.Integer32(p.Identifier)
	}

	s.String(p.formatName(f), p.Description).RawList(p.Raw)

	return s.Format()
}
